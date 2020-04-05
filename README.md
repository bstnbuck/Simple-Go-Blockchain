[![Build Status](https://travis-ci.org/bstnbuck/Simple-Go-Blockchain.svg?branch=master)](https://travis-ci.org/bstnbuck/Simple-Go-Blockchain)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/bstnbuck/Simple-Go-Blockchain/blob/master/LICENSE)
# Simple Go Blockchain

>**Not all functions are implemented yet! This program should only show the principle of blockchains.**

The Blockchain is programmed in Go. As hash algorithm is used SHA512.

## Installation
Nothing to install! It's Go! :D

**Entering some data while executing proof of work now works.**

## Usage
* If you want to execute the Blockchain change to the src directory and execute main.go
* If you want to execute AddTransaction change to the src/addTransaction directory and execute main.go

To execute the program run following commands:

#### Linux:    
  * go build
  * ./src or ./addTransaction

#### Windows:  

  * go build
  * src.exe or addTransaction.exe

Or Debug it in VS-Code or GOLand.

#### Blockchain:
1.  After starting you will be asked how much leading hex-Nulls the hash should have. Enter a decimal like 8, this will be changed to '00000000' <- leading nulls
2.  After that you will be asked how many blocks you want to generate. Please enter a natural number like 0, 1, 25, 199, ...

#### AddTransaction:
1. After starting you will be asked which mode you want to execute. 1 = generate automatically new transactions 2 = manually
2. The rest is self-explanatory.

That's all, depending on the difficulty (leading hex nulls) it may take a while until an output follows.

### Information
* The Code is self-explanatory commented.
* Because of the so named "Nonce" which can be very large the type is Big Integer. This means the "Nonce" can be up to theoretical infinity bits long. This should be sufficient :D
* The payload is filled with random strings of equal length for each block.
* The Proof-of-Work function uses a random string with a incremented Nonce (type bigInt) as hashed operators.

### The following is still being implemented
* Maybe Kotlin or V version
