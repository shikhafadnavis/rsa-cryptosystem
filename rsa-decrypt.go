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

func rsaDecrypt(num *big.Int, N *big.Int, d *big.Int) *big.Int{

	res := squareAndMultiplyWithMod(num, num, d, N)
	return res

}

func main(){

	privateKeyFile := os.Args[1]
	cipherStr := os.Args[2]
	cipher := big.NewInt(0)
	cipher.SetString(cipherStr, 10)
	privKeyByte, readErr := ioutil.ReadFile(privateKeyFile)
	if readErr != nil{
		panic(readErr)
	}
	
	privKeyStr := strings.Split(string(privKeyByte), ",")
	recoveredN := privKeyStr[0]
	recoveredD := privKeyStr[1]
	

	recoveredNInt := big.NewInt(0)
	recoveredDInt := big.NewInt(0)

	recoveredNInt.SetString(recoveredN,10)	
	recoveredDInt.SetString(recoveredD,10)
	
	recovered_plain := rsaDecrypt(cipher, recoveredNInt, recoveredDInt)
	fmt.Println("Recovered Plaintext is: ",recovered_plain)


}
