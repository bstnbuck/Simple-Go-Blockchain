package main

import (
	"bufio"
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

/*
	AddTransaction
	*generates automatically ... new random transactions
	*generates manually ... new transactions
	*saves transaction in payload.gop file

*/

func main() {
	//if file doesn't exist create it
	filename := "src/addTransaction/payload.gop"
	if _, err := os.Stat(filename); err != nil {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
		}

		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}

	//while loop as long as you want to stop it
	stop := false

	for stop == false {
		var mode, times, again int
		fmt.Println("Starting... Generate transactions automatically or manually? [1] = automatically, [2] = manually, [99] = Stop")
		_, err := fmt.Scan(&mode)
		if err != nil {
			fmt.Println("Error in Input: ", err)
		}
		//mode = 1 -> generates automatically new transactions with random users
		if mode == 1 {
			fmt.Println("[AUTO] How many Transactions should be generated? ")
			_, err = fmt.Scan(&times)
			if err != nil {
				fmt.Println("Error in Input: ", err)
			}

			err = runAuto(times, filename)
			if err != nil {
				fmt.Println("Error in Input: ", err)
			}

			fmt.Println("[AUTO] Finished! Any more? [1] = yes [2] = no")
			_, err = fmt.Scan(&again)
			if err != nil {
				fmt.Println("Error in Input: ", err)
			}
			if again == 1 {
				continue
			}

			stop = true

			//mode = 1 -> generates manually new transactions with your input
		} else if mode == 2 {
			fmt.Println("[MANUALLY] How many Transactions should be generated? ")
			_, err = fmt.Scan(&times)
			if err != nil {
				fmt.Println("Error in Input: ", err)
			}

			err = runMan(times, filename)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("[MANUALLY] Finished! Any more? [1] = yes [2] = no")
			_, err = fmt.Scan(&again)
			if err != nil {
				fmt.Println("Error in Input: ", err)
			}
			if again == 1 {
				continue
			}

			stop = true

			//when you like to stop it
		} else if mode == 99 {
			fmt.Println("Stopping...")
			stop = true
		} else {
			fmt.Println("Wrong entry, mistyped?")
		}
	}
	fmt.Println("Press ENTER to continue... ")
	_, err := fmt.Scanln()
	//_ ,err = fmt.Scanln()		//only for windows
	if err != nil {
		fmt.Println(err)
	}
}

//for manually generation
func runMan(times int, filename string) error {
	//open file with write privileges
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	var amount float32
	//make a new string array with "times" sets
	transactions := make([]string, times)
	i := 0
	//ask as long as you like to enter new transactions
	for i < times {
		i++
		fmt.Println("Next Transaction number: ", i)
		fmt.Print("From: ")
		//because of differences between Windows (\r\n) and Linux (\n) decide the correct
		buffer := bufio.NewReader(os.Stdin)
		//from, err := buffer.ReadString('\r')		//only for windows
		from, err := buffer.ReadString('\n') //for linux
		if err != nil {
			return err
		}

		fmt.Print("To: ")
		//to, err := buffer.ReadString('\r')		//only for windows
		to, err := buffer.ReadString('\n') //for linux
		if err != nil {
			return err
		}

		fmt.Print("Amount: ")
		_, err = fmt.Scan(&amount)
		if err != nil {
			return err
		}

		//adjustments for Windows and Linux
		//transactions[i-1] = fmt.Sprintf("New Transaction Timestamp:%v\n From:%s To:%s Transfer:%f\n", time.Now(), from[1:len(from)-1], to[1:], amount) 	//only for windows
		transactions[i-1] = fmt.Sprintf("New Transaction Timestamp:%v\n From:%s To:%s Transfer:%f\n", time.Now(), from[:len(from)-1], to, amount) //for linux
	}
	//add the entered transactions to file
	err = addTransaction(file, transactions)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

//generates automatically new transactions
func runAuto(times int, filename string) error {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	//generate random transactions
	transactions := generateTransactions(times)
	//append them to the file
	err = addTransaction(file, transactions)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

//make random transactions
func generateTransactions(count int) []string {
	//var from, to string
	output := make([]string, count)
	i := 0
	for i < count {
		i++
		output[i-1] = fmt.Sprintf("New Transaction Timestamp:%v\n From:%s To:%s\n Transfer:%f\n", time.Now(), GenerateRandomString(), GenerateRandomString(), GenerateRandomFloat())
	}
	return output
}

//adds transactions to file
func addTransaction(file io.Writer, transaction []string) error {
	writer := bufio.NewWriter(file)

	//iterate trough transaction array and move them to file
	for _, output := range transaction {
		_, err := fmt.Fprintf(writer, "%v", output)
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}
	}
	return nil
}

//generates some random floats
func GenerateRandomFloat() float32 {
	rand.Seed(time.Now().UnixNano())
	random := rand.Float32()

	return random
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString() string {
	b := GenerateRandomBytes(12)
	return base64.URLEncoding.EncodeToString(b) //encode random byte array to base64 encoding
}

func GenerateRandomBytes(n int) []byte {
	b := make([]byte, n)    //new byte array of length n
	_, err := crand.Read(b) //fill array with random
	if err != nil {         //if error print
		println(err)
	}
	return b //return array
}
