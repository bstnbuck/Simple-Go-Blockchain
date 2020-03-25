package main

import (
	"bufio"
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"time"
)

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString() string {
	b := GenerateRandomBytes(12)
	return base64.URLEncoding.EncodeToString(b)		//encode random byte array to base64 encoding
}

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)		//new byte array of length n
	_, err := crand.Read(b)		//fill array with random
	if err != nil {				//if error print
		println(err)
	}
	return b					//return array
}

//calculate hashrate of last created block
func calculateHashrate(timestamp time.Duration, count uint64)string{
	if timestamp.Milliseconds() >= 1{
		//calculate using count of hashes divided by block generation duration
		hashrate := count/uint64(timestamp.Milliseconds())
		//fmt.Println(uint64(timestamp.Milliseconds()))
		//fmt.Println(count)
		returnHashrate := "Hashrate: "+strconv.FormatUint(hashrate, 10)+" H/ms (Hashes per millisecond)"
		return returnHashrate
	}else{
		return "Hashrate: "+strconv.FormatUint(count, 10)+" H/ms (Hashes per millisecond)"
	}
}

//write every block to a .txt file
func writeBlockToBlockchainFile(file io.Writer, output string) error{

	writer := bufio.NewWriter(file)

	//output is same as readable output in makeBlock function
	_, err := fmt.Fprintf(writer, "%v\n\n", output)
	err = writer.Flush()
	if err != nil{
		return err
	}
	return nil
}

//make a string of leading hex nulls. Input is the entered integer number by user
func makeStringNulls(nulls int) (strnulls string){
	for i := 0; i<nulls; i++{
		strnulls+="0"
	}
	return strnulls
}
