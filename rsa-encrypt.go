/*

        Author: Shikha Fadnavis
        Program : RSA key generation
        Date: 10/29/2017

*/

package main

import(
        "fmt"
        //crypt "crypto/rand"
        "math/big"
        "io/ioutil"
	"strings"
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

func rsaEncrypt(num *big.Int, N *big.Int, e *big.Int) *big.Int{

	res := squareAndMultiplyWithMod(num, num, e, N)
	return res

}

func main(){

	publicKeyFile := os.Args[1]
	plaintextStr := os.Args[2]
	plaintext := big.NewInt(0)
	plaintext.SetString(plaintextStr,10)
	pubKeyByte, readErr := ioutil.ReadFile(publicKeyFile)
	if readErr != nil{
		panic(readErr)
	}
	
	pubKeyStr := strings.Split(string(pubKeyByte), ",")
	recoveredN := pubKeyStr[0]
	recoveredE := pubKeyStr[1]

	recoveredNInt := big.NewInt(0)
	recoveredEInt := big.NewInt(0)

	recoveredNInt.SetString(recoveredN,10)	
	recoveredEInt.SetString(recoveredE,10)
 
	cipher := rsaEncrypt(plaintext, recoveredNInt, recoveredEInt)
	fmt.Println("Encrypted text: ", cipher)
	

}
