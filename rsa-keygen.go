/*

	Author: Shikha Fadnavis
	Program : RSA key generation
	Date: 10/24/2017

*/

package main

import(
	"fmt"
//	"crypto/rand"
	"math/big"
//	"io"
	"math/rand"
)


func squareAndMultiplyWithMod(num *big.Int, exp *big.Int, modn *big.Int) (*big.Int){

	var i int 
	var res, finalRes, preFinalRes *big.Int
	//binExp := strconv.FormatInt(exp,2)
	//fmt.Println(binExp)

	//fmt.Println(reflect.TypeOf(binExp))
	//fmt.Println(len(binExp))
	//Start square and multiply
	binExp := fmt.Sprintf("%b", exp)
	res = big.NewInt(num.Int64())
	finalRes = big.NewInt(0)
	preFinalRes = big.NewInt(0)
	if exp == big.NewInt(1){
		return num
	}
	for i = 1; i < len(binExp); i++{
		if binExp[i] == 49{
			//sq and mul
			res.Mul(res,res)
			preFinalRes.Mod(res,modn)
			res.Mul(preFinalRes,num)
			finalRes.Mod(res,modn)
			fmt.Println("Result so far: ", finalRes)
			
		}else{
			//only sq
			res.Mul(res,res)
			finalRes.Mod(res,num)
			fmt.Println("Result so far: ", finalRes)

		}
	}
	fmt.Println("Returning sq and mul with mod result: ", finalRes)
	return finalRes


}

func squareAndMultiply(num *big.Int, exp *big.Int) (*big.Int){

	var i int 
	var res *big.Int
	//binExp := strconv.FormatInt(exp,2)
	//fmt.Println(binExp)

	//fmt.Println(reflect.TypeOf(binExp))
	//fmt.Println(len(binExp))
	//Start square and multiply
	binExp := fmt.Sprintf("%b", exp)
	res = big.NewInt(num.Int64())
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
	nextb := big.NewInt(0)
	var i *big.Int


	randomNum1 := rand.Int63n((big.NewInt(0).Sub(num,big.NewInt(4))).Int64()) + 2
	fmt.Printf("\n Random Number chosen is: %d", randomNum1)
	//Change this to crypto/rand
	randomNum := big.NewInt(randomNum1)
	//b.Mod(squareAndMultiply(randomNum, factor), num)
	b = squareAndMultiplyWithMod(randomNum, factor, num)
	fmt.Println("\nValue of b0 is: ", b)
	if b.Cmp(big.NewInt(1)) == 0 || b.Cmp(big.NewInt(0).Sub(num,big.NewInt(1))) == 0{
	
		return true
	}

	for i = big.NewInt(0); i.Cmp(pow) == -1; i.Add(i,big.NewInt(1)){
        	mulRes.Mul(b,b)
		nextb.Mod(mulRes, num)
                fmt.Println("\nValue is: ", nextb)
                if nextb.Cmp(big.NewInt(1)) == 0{
			
			return false
		}
		if nextb.Cmp(big.NewInt(0).Sub(num,big.NewInt(1))) == 0{
			
			return true
		}

                b = big.NewInt(nextb.Int64())

        }// end of squaring for
 
				

	return false	

}

func millerRabinPrime(num *big.Int) bool{

	var prevRes, factor *big.Int //float64
	

	a := big.NewInt(0)	
	a.Sub(num,big.NewInt(1))
	fmt.Println("Num minus 1 is: ", a)
	k := big.NewInt(1)
	res := big.NewInt(0)
	modulus := big.NewInt(0)
	for true{
		prevRes = big.NewInt(res.Int64())
		div := squareAndMultiply(big.NewInt(2), k)
		fmt.Println("Num will be divided by: ", div)
		res.Div(a,div)
		fmt.Println("resulting number is: ", res)
		fmt.Println("resulting remainder is: ", modulus.Mod(a,div))
		modulus.Mod(a,div)
		if modulus.Cmp(big.NewInt(0)) == 0{
			k.Add(k,big.NewInt(1))
			fmt.Println("k is now; ", k)
		}else{
			break
		}
	}

	factor = big.NewInt(prevRes.Int64())
	pow := big.NewInt(0)
	pow.Sub(k,big.NewInt(1))

	fmt.Println("the two numbers are: ", factor, pow)

	for j := 0; j < 5; j++{
	
		//Call miller test
		if millerTest(num, factor, pow) == false{
			return false
		}
	
	}// end of randomnums for

	return true

} // end of func

func main(){
/*	var Reader io.Reader
	randNum, err := rand.Int(Reader, big.NewInt(20000))
	if err != nil{
		panic(err)
	}
*/
	//randNum := big.NewInt(224737)
	randNum := big.NewInt(0)
	randNum.SetString("2760727302517",10)

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
		}else{
			fmt.Println("\n Composite Number")
		}

	}
}



