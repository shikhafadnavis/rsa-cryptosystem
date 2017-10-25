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
	"reflect"
	"math"
)


func squareAndMultiply(num int, exp int64) int{

	var i, res int
	binExp := strconv.FormatInt(exp,2)
	fmt.Println(binExp)

	fmt.Println(reflect.TypeOf(binExp))
	fmt.Println(len(binExp))
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

	var a, k, res, prevRes float64
	var b, nextb int64
	var result bool = false
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

	fmt.Printf("the two numbers are %f and %f", prevRes, k-1)

	b = int64(math.Exp2(prevRes)) % int64(num)
	fmt.Printf("\nvalue is: %d", b)
	if b == 1{
		//fmt.Println("Composite")
		result =  false
		return result
	}else if b == -1{
		//fmt.Println("Prime")
		result =  true
		return result
	}

	for i := 0; i < int(k-1); i++{
		nextb = (b*b) % int64(num)
		fmt.Printf("\nValue is: %d", nextb)
		if nextb == int64(num - 1){
			result = true
			break
		}
		if nextb == 1{
			result = false
			break
        	}else if nextb == -1{
                	result = true
			break
        	}

		b = nextb

	}
	
	return result

}

func main(){

	expRes := squareAndMultiply(2,1)
	fmt.Println(expRes)
	//fmt.Printf("%d raised to the power of %d is: %d",expRes,)
	primeRes := millerRabinPrime(11)
	if primeRes == true{
		fmt.Println("\n Prime number")
	}else{
		fmt.Println("\n Composite Number")
	}
}



