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
	"strconv"
//	"reflect"
	"math"
	"math/rand"
)


func squareAndMultiply(num int, exp int64) int{

	var i, res int
	binExp := strconv.FormatInt(exp,2)
	//fmt.Println(binExp)

	//fmt.Println(reflect.TypeOf(binExp))
	//fmt.Println(len(binExp))
	//Start square and multiply
	
	res = num
	if exp == 1{
		return num
	}
	for i = 1; i < len(binExp); i++{
		if binExp[i] == 49{
			//sq and mul
			res = res * res
			res = res * num
		}else{
			//only sq
			res = res * res

		}
	}

	return res


}

func millerRabinPrime(num float64) bool{

	var a, k, res, prevRes, factor, pow float64
	var b, nextb int64
	var result bool = true
	a = num - 1
	k = 1
	for true{
		prevRes = res
		res = a/math.Exp2(k)
		if res == float64(int64(res)){
			k++
		}else{
			break
		}
	}

	factor = prevRes
	pow = k-1

	fmt.Printf("the two numbers are %f and %f", factor, pow)

	for j := 0; j < 5; j++{
	
		randomNum := rand.Int63n(int64(num-4)) + 2
		fmt.Printf("\n Random Number chosen is: %d", randomNum)
		//randomNum := 2
		b = int64(squareAndMultiply(int(randomNum), int64(factor))) % int64(num)
		//b = int64(math.Exp2(factor)) % int64(num)
		fmt.Println("\nValue of b0 is: ", b)
		if b == 1 || b == int64(num-1){
			result = true
			return result
		}
		//fmt.Println("Composite")
		for i := 0; i < int(k-1); i++{
                        nextb = (b*b) % int64(num)
                       	fmt.Println("\nValue is: ", nextb)
                       	if nextb == 1{
				result = false
				return result
			}
			if nextb == int64(num-1){
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
	primeRes := millerRabinPrime(281)
	if primeRes == true{
		fmt.Println("\n Prime number")
	}else{
		fmt.Println("\n Composite Number")
	}
}



