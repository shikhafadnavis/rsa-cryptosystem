/*

	Author: Shikha Fadnavis
	Program : RSA key generation
	Date: 10/24/2017

*/

package main

import(
	"fmt"
//	"crypto/rand"
//	"strings"
//	"strconv"
//	"reflect"
//	"math"
	"math/big"
//	"math/rand"
)


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
			//res = res * res
			res.Mul(res,res)
			//res = res * num
			res.Mul(res,num)
		}else{
			//only sq
			//res = res * res
			res.Mul(res,res)

		}
	}

	return res


}

func millerRabinPrime(num *big.Int) bool{

	var prevRes, factor *big.Int //float64
	var i *big.Int //int64
	var result bool = true
	a := big.NewInt(0)	
	a.Sub(num,big.NewInt(1))
	fmt.Println("Num minus 1 is: ", a)
	k := big.NewInt(1)
	res := big.NewInt(0)
	modulus := big.NewInt(0)
	b := big.NewInt(0)
	mulRes := big.NewInt(0)
	nextb := big.NewInt(0)
	for true{
		prevRes = big.NewInt(res.Int64())
		div := squareAndMultiply(big.NewInt(2), k)
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
	
		//randomNum := rand.Int63n(int64(num-4)) + 2
		//fmt.Printf("\n Random Number chosen is: %d", randomNum)
		randomNum := big.NewInt(2)
		b.Mod(squareAndMultiply(randomNum, factor), num)
		//b = int64(squareAndMultiply(int(randomNum), int64(factor))) % int64(num)
		//b = int64(math.Exp2(factor)) % int64(num)
		fmt.Println("\nValue of b0 is: ", b)
		if b.Cmp(big.NewInt(1)) == 0 || b.Cmp(big.NewInt(0).Sub(num,big.NewInt(1))) == 0{
			result = true
			return result
		}

		for i = big.NewInt(0); i.Cmp(pow) == -1; i.Add(i,big.NewInt(1)){
                        //nextb = (b*b) % int64(num)
			mulRes.Mul(b,b)
			nextb.Mod(mulRes, num)
                       	fmt.Println("\nValue is: ", nextb)
                       	if nextb.Cmp(big.NewInt(1)) == 0{
				result = false
				return result
			}
			if nextb.Cmp(big.NewInt(0).Sub(num,big.NewInt(1))) == 0{
				result = true
				return result
			}

                       	b = big.NewInt(nextb.Int64())

                }// end of squaring for
 
				

		result = false
		return result	
	
	}// end of randomnums for

	result = true
	return result

} // end of func

func main(){

	//expRes := squareAndMultiply(2,1)
	//fmt.Println(expRes)
	//fmt.Printf("%d raised to the power of %d is: %d",expRes,)
	primeRes := millerRabinPrime(big.NewInt(1223))
	if primeRes == true{
		fmt.Println("\n Prime number")
	}else{
		fmt.Println("\n Composite Number")
	}
}



