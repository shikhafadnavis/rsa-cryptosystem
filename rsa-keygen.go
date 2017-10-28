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
	res = num
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

	var a, k, res, prevRes, factor, pow, modulus, mulRes *big.Int //float64
	var b, nextb, i *big.Int //int64
	var result bool = true
	
	a.Sub(num,big.NewInt(1))
	fmt.Println("Num minus 1 is: ", a)
	k = big.NewInt(1)
	for true{
		prevRes = res
		div := squareAndMultiply(big.NewInt(2), k)
		res.Div(a,div)
		if modulus.Mod(a,div) == big.NewInt(0){
			k.Add(k,big.NewInt(1))
		}else{
			break
		}
	}

	factor = prevRes
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
		if b == big.NewInt(1) || b == big.NewInt(0).Sub(num,big.NewInt(1)){
			result = true
			return result
		}

		for i = big.NewInt(0); i.Cmp(pow) == -1; i.Add(i,big.NewInt(1)){
                        //nextb = (b*b) % int64(num)
			mulRes.Mul(b,b)
			nextb.Mod(mulRes, num)
                       	fmt.Println("\nValue is: ", nextb)
                       	if nextb == big.NewInt(1){
				result = false
				return result
			}
			if nextb == big.NewInt(0).Sub(num,big.NewInt(1)){
				result = true
				return result
			}

                       	b = nextb

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
	primeRes := millerRabinPrime(big.NewInt(85))
	if primeRes == true{
		fmt.Println("\n Prime number")
	}else{
		fmt.Println("\n Composite Number")
	}
}



