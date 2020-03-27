[![Build Status](https://travis-ci.com/bstnbuck/Simple-Go-Blockchain.svg?branch=master)](https://travis-ci.com/bstnbuck/Simple-Go-Blockchain)
# Simple Go Blockchain

>**Not all functions are implemented yet! This program should only show the principle of blockchains.**

The Blockchain is programmed in Go. As hash algorithm is used SHA512.

## Installation
Nothing to install! It's Go! :D

## Usage
1.  After starting you will be asked how much leading hex-Nulls the hash should have. Enter a decimal like 8, this will be changed to '00000000' <- leading nulls
2.  After that you will be asked how many blocks you want to generate. Please enter a natural number like 0, 1, 25, 199, ...

That's all, depending on the difficulty (leading hex nulls) it may take a while until an output follows.

### Information
* The Code is self-explanatory commented.
* Because of the so named "Nonce" which can be very large the type is Big Integer. This means the "Nonce" can be up to theoretical infinity bits long. This should be sufficient :D
* The payload is filled with random strings of equal length for each block.
* The Proof-of-Work function uses a random string with a incremented Nonce (type bigInt) as hashed operators.

### The following is still being implemented
* Entering some data while executing proof of work.
* Maybe Kotlin version
