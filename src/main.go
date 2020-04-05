package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Simple Go Blockchain!")
	//generate genesis block and append it to the blockchain
	genesisBlock := makeGenesisBlock()
	Blockchain = append(Blockchain,genesisBlock)

	var times, nulls int
	fmt.Println("Starting... How much leading hex nulls? ")
	_, err:= fmt.Scan(&nulls)
	if err != nil{
		fmt.Println("Error in Input: ",err)
	}

	fmt.Println("How many blocks should be generated? ")
	_, err= fmt.Scan(&times)
	if err != nil{
		fmt.Println("Error in Input: ",err)
	}

	fmt.Println("Started at: ",time.Now())
	err = run(times, nulls)
	if err != nil{
		fmt.Println("Error in RUN: ",err)
	}

	fmt.Println("All Blocks generated! Press ENTER to continue... ")
	_ ,err = fmt.Scanln()
	//_ ,err = fmt.Scanln()		//only for windows
	if err != nil {
		fmt.Println(err)
	}
}
