package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strconv"
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

	/* removed at 05/04/20, instead use real transactions
	//Define Payload, a bit more than one string please;)
	payload := ""; i := 0
	for i <= 30{
		payload += " "+GenerateRandomString()//RandStringRunes()
		if i == 15{
			payload += "\n"
		}
		i++
	}
	*/
	block.payload = getTransactions()

	//make readable output
	output := fmt.Sprintf("New Block Index:%v Timestamp:%v \nHashPoW:%v \nText&Nonce:%v \nPrevHashHeader:%v \nHashHeader:%v \nData:\n%v",
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
