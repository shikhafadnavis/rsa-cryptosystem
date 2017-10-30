/*

	Author: Shikha Fadnavis
	Program : RSA key generation
	Date: 10/24/2017

*/

package main

import(
	"fmt"
	crypt "crypto/rand"
	"math/big"
	"io/ioutil"
	"os"
)


func squareAndMultiplyWithMod(a *big.Int, a2 *big.Int, b *big.Int, c *big.Int) (*big.Int){

	var i int 
	var startVal, res, someRes, preRes *big.Int

	binExp := fmt.Sprintf("%b", b)
	if b == big.NewInt(1){
		return a
	}

	// Retain original value
	startVal = big.NewInt(0)
	startVal.Mod(a,c)

	res = big.NewInt(0)
	res.Mod(a2,c)

	for i = 1; i < len(binExp); i++{	
		// Square regardless
		someRes = big.NewInt(0)
		someRes.Mul(res, res)
		res.Mod(someRes, c)

		if binExp[i] == 49{
			preRes = big.NewInt(0)
			preRes.Mul(res, startVal)
			res.Mod(preRes,c)			
		}

	}

	return res	
		


}

func squareAndMultiply(num *big.Int, exp *big.Int) (*big.Int){

	var i int 
	var res *big.Int
	//Start square and multiply
	binExp := fmt.Sprintf("%b", exp)
	if exp == big.NewInt(1){
		return num
	}
	for i = 1; i < len(binExp); i++{
		if binExp[i] == 49{
			//sq and mul
			res.Mul(res,res)
			res.Mul(res,num)
			
		}else{
			//only sq
			res.Mul(res,res)

		}
	}

	return res

}


func millerTest(num *big.Int, factor *big.Int, pow *big.Int) bool{


	b := big.NewInt(0)
	mulRes := big.NewInt(0)
	var i *big.Int

	randNumMiller := big.NewInt(0)
	randNumByteMiller := make([]byte,64)
        crypt.Read(randNumByteMiller)
        randNumMiller.SetBytes(randNumByteMiller)

	b = squareAndMultiplyWithMod(randNumMiller, randNumMiller, factor, num)

	if b.Cmp(big.NewInt(1)) == 0 || b.Cmp(big.NewInt(0).Sub(num,big.NewInt(1))) == 0{
	
		return true
	}

	for i = big.NewInt(0); i.Cmp(pow) == -1; i.Add(i,big.NewInt(1)){
        	mulRes.Mul(b,b)
		b.Mod(mulRes, num)
                if b.Cmp(big.NewInt(1)) == 0{
			
			return false
		}
		if b.Cmp(big.NewInt(0).Sub(num,big.NewInt(1))) == 0{
			
			return true
		}

        }// end of squaring for
 
				

	return false	

}

func millerRabinPrime(num *big.Int) bool{

	var factor *big.Int 
	

	a := big.NewInt(0)	
	a.Sub(num,big.NewInt(1))
	k := big.NewInt(0) 
	modulus := big.NewInt(0)
	for true{
		a.Div(a,big.NewInt(2)) 
		modulus.Mod(a,big.NewInt(2))
		if modulus.Cmp(big.NewInt(0)) == 0{
			k.Add(k,big.NewInt(1))
		}else{
			break
		}
	}

	factor = big.NewInt(0)
	factor.Mul(a,big.NewInt(2))
	pow := big.NewInt(0)
	pow.Sub(k,big.NewInt(0)) 

	for j := 0; j < 5; j++{
	
		//Call miller test
		if millerTest(num, factor, pow) == false{
			return false
		}
	
	}// end of randomnums for

	return true

} // end of func


func randGenerate() *big.Int{
	
	randNum := big.NewInt(0)
	
	for true{
		randNumByte := make([]byte,64)
		crypt.Read(randNumByte)
		randNum.SetBytes(randNumByte)
		fmt.Println("random number chosen is: ", randNum)
		operation := big.NewInt(0)
		operation.Mod(randNum, big.NewInt(2))
		if operation.Cmp(big.NewInt(0)) == 0{
			fmt.Println("Composite Number")
			//generate random again
		}else{
			primeRes := millerRabinPrime(randNum)
			if primeRes == true{
				fmt.Println("\n Prime number")
				break
			}else{
				fmt.Println("\n Composite Number")
			}
		}
	}

	return randNum

}

func extendedEucledian(a *big.Int, b *big.Int) (*big.Int, *big.Int, *big.Int){

	d := big.NewInt(0)
	x := big.NewInt(0)
	y := big.NewInt(0)
	x1 := big.NewInt(0)
	x2 := big.NewInt(1)
	y1 := big.NewInt(1)
	y2 := big.NewInt(0)
	q := big.NewInt(0)
	r := big.NewInt(0)
	if b.Cmp(big.NewInt(0)) == 0{
		d.Set(a)
		x = big.NewInt(1)
		y = big.NewInt(0)
		return d,x,y
	}
		for b.Cmp(big.NewInt(0)) == 1{
			q.Div(a,b)
			r.Sub(a, big.NewInt(0).Mul(q,b))
			x.Sub(x2, big.NewInt(0).Mul(q,x1))
			y.Sub(y2, big.NewInt(0).Mul(q,y1))
			
			a.Set(b)
			b.Set(r)
			x2.Set(x1)
			x1.Set(x)
			y2.Set(y1)
			y1.Set(y)
			
		}
		


	d.Set(a)
	x.Set(x2)
	y.Set(y2)


	return d,x,y
}





func writePubKey(N *big.Int, e *big.Int, filename string){
	NStr := N.String()
	comma := ","
	eStr := e.String()
	eFileStr := NStr + comma + eStr
	eFileByte := []byte(eFileStr)
	writeErr := ioutil.WriteFile(filename,eFileByte, 0644)
	if writeErr != nil{
		panic(writeErr)
	}	
}

func writePrivKey(N *big.Int, d *big.Int, p *big.Int, q *big.Int, filename string){
	NStr := N.String()
	comma := ","
	dStr := d.String()
	pStr := p.String()
	qStr := q.String()

	dFileStr := NStr + comma + dStr + comma + pStr + comma + qStr
	dFileByte := []byte(dFileStr)
	writeErr := ioutil.WriteFile(filename, dFileByte, 0644)
	if writeErr != nil{
		panic(writeErr)
	}	
	


}

func main(){

	modulus := big.NewInt(0)
	prime1minus := big.NewInt(0)
	prime2minus := big.NewInt(0)	
	phi := big.NewInt(0)
	privExp := big.NewInt(0)

	prime1 := randGenerate()
	prime2 := randGenerate()
	
	modulus.Mul(prime1, prime2)

	prime1minus.Sub(prime1, big.NewInt(1))
	prime2minus.Sub(prime2, big.NewInt(1))

	phi.Mul(prime1minus, prime2minus)

	pubExp := randGenerate()

	fmt.Println("Prime 1 is: ", prime1)
        fmt.Println("Prime 2 is: ", prime2)
        fmt.Println("Public Modulus is: ", modulus)
        fmt.Println("Phi Modulus is: ", phi)

	fmt.Println("Public Exponent is: ", pubExp)

	//Insert check for coprimality with phi

	pubExpDup := big.NewInt(0)
	pubExpDup.Set(pubExp)
	phiDup := big.NewInt(0)
	phiDup.Set(phi)
	val1, val2, val3 := extendedEucledian(pubExpDup, phiDup)

	fmt.Println("Phi Modulus after extended euc is: ", phi)

	fmt.Println("val1 is: ", val1)
	fmt.Println("val2 is: ", val2)
	fmt.Println("val3 is: ", val3)

	if val2.Cmp(big.NewInt(0)) == -1{
		val2.Add(val2, phi)
	}

	fmt.Println("Private exponent is: ", val2)
	privExp.Set(val2)	

	fmt.Println("PubExp*val2 = ", big.NewInt(0).Mul(pubExp, val2))

	pubWithVal2 := big.NewInt(0).Mul(pubExp, val2)
	pubWithVal2.Mod(pubWithVal2, phi)
	fmt.Println("Pub with Val2: ", pubWithVal2)

	// File operations
	
	writePubKey(modulus, pubExp, os.Args[1])

	writePrivKey(modulus, privExp, prime1, prime2, os.Args[2])

	

}



