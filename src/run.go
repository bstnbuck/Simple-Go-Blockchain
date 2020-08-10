package main

import (
	"fmt"
	"os"
	"time"
)

func run(times, nulls int) error {

	strnulls := makeStringNulls(nulls)
	//create output file

	filename := "src/examples/blckchn_40bit.txt" //change your working directory if it doesn't work
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
