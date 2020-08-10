package main

import (
	"crypto/sha512"
	"encoding/hex"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

//proof of work function
func pow(text string, nulls string) (hasht string, textNonce string, count uint64) {
	//the nonce can be very large, therefore big int -> This means the nonce can be up to theoretical infinity bits long. This should be sufficient. :D
	rand.Seed(time.Now().UnixNano())
	var nonce = big.NewInt(rand.Int63n(1000000000)) //edit 24.03.
	//var nonce = uint64(rand.Int63n(1000000000))		//nonce can be theoretical up to 512 bit, uint64 only 64bit, therefore big integer -> edit 24.03.

	count = 0
	//endless for loop until hash is found
	hash := sha512.New() //hash initialization before hashing -> edit 23.03.20
	for {
		//hash := sha512.New() 	//hash initialization before hashing -> edit 23.03.20
		hash.Write([]byte(text + nonce.String()))
		//hash.Write([]byte(text+strconv.FormatUint(nonce, 10)))
		hasht := hex.EncodeToString(hash.Sum(nil))
		//fmt.Println(hasht)

		//check if the hash has the required leading nulls
		if strings.HasPrefix(hasht, nulls) {
			//fmt.Println("Hash found! ",hasht)
			//fmt.Println("Text and Nonce: ",text," + ",nonce)
			return hasht, text + nonce.String(), count //edit 24.03.
			//return hasht, text+strconv.FormatUint(nonce, 10), count		//edit 24.03.
		}
		hash.Reset()
		//if not, nonce will be incremented
		nonce.Add(nonce, big.NewInt(1)) //edit 24.03.
		//nonce++								//edit 24.03.
		count++
	}
}
