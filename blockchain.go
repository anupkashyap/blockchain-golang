package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math"
	"strconv"
	"time"
)

type Block map[string]string

type Blockchain struct {
	chain []Block
}

//Initializes the blockchain with the genesis block
func (bc *Blockchain) initialize() {
	//Genesis block
	bc.createBlock(1, "0")
}

// Created and adds a new block to the chain
func (bc *Blockchain) createBlock(proof int, previousHash string) Block {
	block := make(Block)
	block["index"] = strconv.Itoa(len(bc.chain) + 1)
	block["timestamp"] = time.Now().String()
	block["proof"] = strconv.Itoa(proof)
	block["previousHash"] = previousHash
	bc.chain = append(bc.chain, block)
	return block
}

//Gets the latest block form the chain
func (bc Blockchain) getPreviousBlock() Block {
	return bc.chain[len(bc.chain)-1]
}

//Calculates the proof of work
func (bc Blockchain) proofOfWork(previousProof int) int {
	newProof := 1
	checkProof := false
	for !checkProof {
		//Hash Operation
		hashBytes := sha256.Sum256([]byte(strconv.Itoa(int(math.Pow(float64(newProof), 2) - math.Pow(float64(previousProof), 2)))))
		hash := hex.EncodeToString(hashBytes[:])
		if string(hash[:4]) == "0000" {
			checkProof = true
		} else {
			newProof++
		}

	}
	return newProof
}

// Creates the hash of a block
func (bc Blockchain) hash(b Block) string {
	encodedBlock, _ := json.Marshal(b)
	hashBytes := sha256.Sum256(encodedBlock)
	return hex.EncodeToString(hashBytes[:])
}

//Checks the integrity of the blockchain
func (bc Blockchain) isChainValid(chain []Block) bool {
	prevBlock := chain[0]
	blockIndex := 1
	for blockIndex < len(chain) {
		block := chain[blockIndex]
		if block["previousHash"] != bc.hash(prevBlock) {
			return false
		}
		previousProof, _ := strconv.Atoi(prevBlock["proof"])
		currentProof, _ := strconv.Atoi(block["proof"])
		//Hash Operation
		hashBytes := sha256.Sum256([]byte(strconv.Itoa(int(math.Pow(float64(currentProof), 2) - math.Pow(float64(previousProof), 2)))))
		hash := hex.EncodeToString(hashBytes[:])
		if string(hash[:4]) != "0000" {
			return false
		}
		prevBlock = block
		blockIndex++
	}
	return true
}
