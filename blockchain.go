package main

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//Define how a Block should look like
type Block struct{
	//Header elements
	Index int64
	Timestamp time.Time
	HashPoW string
	textNoncePoW string
	PrevHashHeader string

	//Hash of the header
	hashHeader string

	//Data
	payload string
}

//Blockchain is a Array of valid Blocks
var Blockchain []Block


func makeBlock() (string, Block){
	//use random text in PoW
	text := RandStringRunes()
	block := Block{}

	//Find suitable hash, get the hash and the used nonce + text back
	block.HashPoW, block.textNoncePoW = pow(text, "000000")

	//Define Header elements
	block.Index = int64(len(Blockchain))
	block.Timestamp = time.Now()
	block.PrevHashHeader = Blockchain[len(Blockchain)-1].hashHeader

	//make hash of Header elements
	block.hashHeader = makeBlockHash(block)

	//Define Payload, a bit more than one string please;)
	payload := ""; i := 0
	for i <= 10{
		payload += " "+RandStringRunes()
		i++
	}
	block.payload = payload		//these normally are transactions!

	//make readable output
	output := fmt.Sprintf("New Block Index:%v Timestamp:%v \nHashPoW:%v \nText&Nonce:%v \nPrevHashHeader:%v \nHashHeader:%v \nData:%v",
		block.Index,block.Timestamp, block.HashPoW, block.textNoncePoW,block.PrevHashHeader,block.hashHeader,block.payload)

	//return the output and block
	return output, block
}

//function to make hash of the block header
func makeBlockHash(block Block) string{
	hash := sha512.New()
	hash.Write([]byte(strconv.FormatInt(block.Index, 10)+block.Timestamp.String()+block.HashPoW+block.textNoncePoW+block.PrevHashHeader))
	hashHeader := hex.EncodeToString(hash.Sum(nil))
	return hashHeader
}

//proof if the new created block is valid
func isNewBlockValid(newBlock Block) bool{
	lastBlock := Blockchain[len(Blockchain)-1]
	if lastBlock.hashHeader == newBlock.PrevHashHeader && lastBlock.Index+1 == newBlock.Index && newBlock.hashHeader == makeBlockHash(newBlock){
		return true
	}else{
		return false
	}
}

//initialise the blockchain with a first block, filled with null elements
func makeGenesisBlock() Block{
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
func pow(text string, nulls string) (hasht string, textNonce string){
	//the nonce can be very large, therefore unsigned int 64bit -> This means the nonce can be up to 2**64 bits long. This should be sufficient.
	var nonce uint64 = uint64(rand.Intn(1000000000))

	//endless for loop until hash is found
	for {
		hash := sha512.New()
		hash.Write([]byte(text+strconv.FormatUint(nonce, 10)))
		hasht := hex.EncodeToString(hash.Sum(nil))
		//fmt.Println(hasht)

		//check if the hash has the required leading nulls
		if strings.HasPrefix(hasht, nulls){
			//fmt.Println("Hash found! ",hasht)
			//fmt.Println("Text and Nonce: ",text," + ",nonce)
			return hasht, text+strconv.FormatUint(nonce, 10)
		}
		hash.Reset()
		//if not nonce will be incremented
		nonce += 1
	}
}

//write every block to a .txt file
func writeBlockToBlockchainFile(file io.Writer, output string) error{

	writer := bufio.NewWriter(file)

	//output is same as readable output in makeBlock function
	_, err := fmt.Fprintf(writer, "%v\n\n", output)
	if err != nil{
		return err
	}
	writer.Flush()
	return nil
}

//make random strings of length 10 for hash and payload
func RandStringRunes() string {
	n := 10
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!=?/(),.+-#")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//automatically adjust the difficulty of PoW
func adjustDifficulty(timestamp time.Duration) string{
	//adjust difficulty of PoW automatically by using timestamp (time to make new block)
	//will be implemented soon
	return ""
}

func run(times int) error{
	//create output file
	file, err := os.Create("blockchain.txt")
	//fmt.Println(err)
	if err != nil{
		return err
	}

	//create ... times new blocks
	i:= 0
	for i != times{
		start := time.Now()

		output, newBlock := makeBlock()
		fmt.Println(output)
		//return output

		//check if new Block is valid
		isValid := isNewBlockValid(newBlock)
		if isValid{
			//if true append it to the actual blockchain
			Blockchain = append(Blockchain,newBlock)
			fmt.Println("Block valid!")
			//and file...
			err := writeBlockToBlockchainFile(file, output)
			//fmt.Println(err)
			if err != nil{
				return err
			}
		}else{
			fmt.Println("Last Block isn't valid, so it will not append to the Blockchain!")
		}

		t := time.Now()
		//elapsed shows how long the block generation took
		elapsed := t.Sub(start)
		outputElapsed := "Time elapsed: "+elapsed.String()
		writeBlockToBlockchainFile(file,outputElapsed)
		fmt.Print("Time to make new Block: ",elapsed,"\n\n")
		_ = adjustDifficulty(elapsed)

		i++
	}
	file.Close()
	return nil
}


func main(){
	fmt.Println("Go Blockchain!")
	//generate genesis block and append it to the blockchain
	genesisBlock := makeGenesisBlock()
	Blockchain = append(Blockchain,genesisBlock)

	var times int
	fmt.Println("Starting... How much Blocks would you like to generate? ")
	_, err:= fmt.Scan(&times)
	if err != nil{
		return
	}
	fmt.Println("Started at: ",time.Now())
	run(times)

	//fmt.Println(Blockchain)
}
