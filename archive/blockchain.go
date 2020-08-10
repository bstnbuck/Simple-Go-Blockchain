package main

import (
	"bufio"
	crand "crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//Block define how a Block should look like
type Block struct {
	//Header elements
	Index          int64
	Timestamp      time.Time
	HashPoW        string
	textNoncePoW   string
	PrevHashHeader string

	//Hash of the header
	hashHeader string

	//Data
	payload string
}

//Blockchain is a Array of valid Blocks
var Blockchain []Block

func makeBlock(nulls string) (string, Block, uint64) {
	//use random text in PoW
	text := GenerateRandomString() //RandStringRunes()
	var count uint64
	block := Block{}

	//Find suitable hash, get the hash and the used nonce + text back
	block.HashPoW, block.textNoncePoW, count = pow(text, nulls)

	//Define Header elements
	block.Index = int64(len(Blockchain))
	block.Timestamp = time.Now()
	block.PrevHashHeader = Blockchain[len(Blockchain)-1].hashHeader

	//make hash of Header elements
	block.hashHeader = makeBlockHash(block)

	//Define Payload, a bit more than one string please;)
	payload := ""
	i := 0
	for i <= 30 {
		payload += " " + GenerateRandomString() //RandStringRunes()
		if i == 15 {
			payload += "\n"
		}
		i++
	}
	block.payload = payload //these normally are transactions!

	//make readable output
	output := fmt.Sprintf("New Block Index:%v Timestamp:%v \nHashPoW:%v \nText&Nonce:%v \nPrevHashHeader:%v \nHashHeader:%v \nData:%v",
		block.Index, block.Timestamp, block.HashPoW, block.textNoncePoW, block.PrevHashHeader, block.hashHeader, block.payload)

	//return the output and block
	return output, block, count
}

//function to make hash of the block header
func makeBlockHash(block Block) string {
	hash := sha512.New()
	hash.Write([]byte(strconv.FormatInt(block.Index, 10) + block.Timestamp.String() + block.HashPoW + block.textNoncePoW + block.PrevHashHeader))
	hashHeader := hex.EncodeToString(hash.Sum(nil))
	return hashHeader
}

//proof if the new created block is valid
func isNewBlockValid(newBlock Block) bool {
	lastBlock := Blockchain[len(Blockchain)-1]
	if lastBlock.hashHeader == newBlock.PrevHashHeader && lastBlock.Index+1 == newBlock.Index && newBlock.hashHeader == makeBlockHash(newBlock) {
		return true
	}
	return false
}

//initialise the blockchain with a first block, filled with null elements
func makeGenesisBlock() Block {
	text := "Welcome to this Go-Blockchain!"
	block := Block{}

	//Find suitable hash
	block.HashPoW, block.textNoncePoW = "0", "0"

	//Define Header elements
	block.Index = 0
	block.Timestamp = time.Now()
	block.PrevHashHeader = "0"

	//make hash of Header elements
	block.hashHeader = makeBlockHash(block)

	//Define Data
	block.payload = text
	return block
}

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

//write every block to a .txt file
func writeBlockToBlockchainFile(file io.Writer, output string) error {

	writer := bufio.NewWriter(file)

	//output is same as readable output in makeBlock function
	_, err := fmt.Fprintf(writer, "%v\n\n", output)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString() string {
	b := GenerateRandomBytes(12)
	return base64.URLEncoding.EncodeToString(b) //encode random byte array to base64 encoding
}

// GenerateRandomBytes creates random bytes for randomString
func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)    //new byte array of length n
	_, err := crand.Read(b) //fill array with random
	if err != nil {         //if error print
		println(err)
	}
	return b //return array
}

//calculate hashrate of last created block
func calculateHashrate(timestamp time.Duration, count uint64) string {
	if timestamp.Milliseconds() >= 1 {
		//calculate using count of hashes divided by block generation duration
		hashrate := count / uint64(timestamp.Milliseconds())
		//fmt.Println(uint64(timestamp.Milliseconds()))
		//fmt.Println(count)
		returnHashrate := "Hashrate: " + strconv.FormatUint(hashrate, 10) + " H/ms (Hashes per millisecond)"
		return returnHashrate
	}
	return "Hashrate: " + strconv.FormatUint(count, 10) + " H/ms (Hashes per millisecond)"
}

//make a string of leading hex nulls. Input is the entered integer number by user
func makeStringNulls(nulls int) (strnulls string) {
	for i := 0; i < nulls; i++ {
		strnulls += "0"
	}
	return strnulls
}

func run(times, nulls int) error {

	strnulls := makeStringNulls(nulls)
	//create output file
	filename := "blckchn_36bit_2.txt"
	file, err := os.Create(filename)
	//fmt.Println(err)
	if err != nil {
		return err
	}
	//close file as long as no new block is generated
	err = file.Close()
	if err != nil {
		return err
	}

	//create ... times new blocks
	i := 0
	for i != times {
		start := time.Now()

		output, newBlock, count := makeBlock(strnulls)
		fmt.Println(output)

		//check if new Block is valid
		isValid := isNewBlockValid(newBlock)
		if isValid {
			//if true append it to the actual blockchain
			Blockchain = append(Blockchain, newBlock)
			fmt.Println("Block valid!")

			//and file...
			file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0755)
			if err != nil {
				return err
			}
			err = writeBlockToBlockchainFile(file, output)
			//fmt.Println(err)
			if err != nil {
				return err
			}

			t := time.Now()
			//elapsed shows how long the block generation took
			elapsed := t.Sub(start)
			outputElapsed := "Time elapsed: " + elapsed.String() + "\n"

			//calculate the Hashrate of this block
			hashRate := calculateHashrate(elapsed, count)
			err = writeBlockToBlockchainFile(file, hashRate)
			if err != nil {
				return err
			}
			err = writeBlockToBlockchainFile(file, outputElapsed)
			if err != nil {
				return err
			}
			fmt.Println("Count: ", count)
			fmt.Println(hashRate)
			fmt.Print("Time to make new Block: ", elapsed, "\n\n")
			i++
			err = file.Close()
			if err != nil {
				return err
			}
		} else {
			fmt.Println("Last Block isn't valid, so it will not append to the Blockchain!")
		}
	}
	return nil
}

func main() {
	fmt.Println("Go Blockchain!")
	//generate genesis block and append it to the blockchain
	genesisBlock := makeGenesisBlock()
	Blockchain = append(Blockchain, genesisBlock)

	var times, nulls int
	fmt.Println("Starting... How much leading hex nulls? ")
	_, err := fmt.Scan(&nulls)
	if err != nil {
		fmt.Println("Error in Input: ", err)
	}

	fmt.Println("How much Blocks would you like to generate? ")
	_, err = fmt.Scan(&times)
	if err != nil {
		fmt.Println("Error in Input: ", err)
	}

	fmt.Println("Started at: ", time.Now())
	err = run(times, nulls)
	if err != nil {
		fmt.Println("Error in RUN: ", err)
	}
	//fmt.Println(Blockchain)
}
