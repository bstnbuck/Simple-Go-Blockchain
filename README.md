# Simple Go Blockchain

>**Not all functions are implemented yet! This program should only show the principle of blockchains.**

The Blockchain is programmed in Go. As hash algorithm is used SHA512.

## Installation
Nothing to install! It's Go! :D

## Usage
1. After starting you will be asked how much Blocks you would like to generate. Please enter a natural number like 0, 1, 2, 25, 199, ...

That's all, depending on the difficulty it may take a while until an output follows.

### Information
* The Code is self-explanatory commented.
* Because of the so named "Nonce" which can be very large the type is uInt64 (unsigned Integer 64bit). This means the "Nonce" can be up to 2**64 bits long. This should be sufficient.
* The payload is filled with random strings of equal length for each block.
* The Proof-of-Work function uses a random string with a incremented Nonce (type uInt64) as hashed operators.

### The following is still being implemented
* AI to automatically adjust the difficulty.
* Entering some data while executing proof of work.
* Kotlin version
