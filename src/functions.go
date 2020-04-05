package main

import (
	"bufio"
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
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

//adds last 10 generated transaction with lifo-principle (last in -> first out) to the payload of new generated block
func getTransactions()string{
	filename := "src/addTransaction/payload.gop"
	//If file doesn't exist, print error and return
	if _, err := os.Stat(filename); err != nil{
		fmt.Println("Error getTransactions: ",err)
		return ""
	}

	//if exist, open file
	file, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
		 fmt.Println("Error getTransactions: ",err)
	}

	scanner := bufio.NewScanner(file)
	counter := 27
	progress := 0
	var output, text string
	//scan first 10 transactions (27 lines) into output string, all others that follow save to another variable
	for scanner.Scan() {
		if counter > progress {
			output += scanner.Text()+"\n"
			progress++
		}else{
			text += scanner.Text()+"\n"
		}
	}
	err = file.Close()
	if err != nil {
		fmt.Println("Error getTransactions: ",err)
	}

	//create file new (clear it) and move all others that follow after the 10 transactions into it
	file, err = os.Create(filename)
	file, err = os.OpenFile(filename, os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("Error getTransactions: ",err)
	}
	writer := bufio.NewWriter(file)
	_,err =file.WriteString(text)
	if err != nil{
		fmt.Println("Error getTransactions: ",err)
	}

	err = writer.Flush()
	if err != nil{
		fmt.Println("Error getTransactions: ",err)
	}

	err = file.Close()
	if err != nil {
		fmt.Println("Error getTransactions: ",err)
	}

	return output
}
