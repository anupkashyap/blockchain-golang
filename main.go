package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	blockchain := Blockchain{}
	blockchain.initialize()

	r := gin.Default()

	//  localhost:8080/mineBlock
	r.GET("/mineBlock", func(c *gin.Context) {
		prevBlock := blockchain.getPreviousBlock()
		prevProof, _ := strconv.Atoi(prevBlock["proof"])
		proof := blockchain.proofOfWork(prevProof)
		prevHash := blockchain.hash(prevBlock)
		block := blockchain.createBlock(proof, prevHash)
		c.JSON(200, gin.H{
			"message":      "Successfully mined a block",
			"index":        block["index"],
			"timestamp":    block["timestamp"],
			"proof":        block["proof"],
			"previousHash": block["previousHash"],
		})
	})

	//  localhost:8080/getChain
	r.GET("/getChain", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"chain":  blockchain.chain,
			"Length": len(blockchain.chain),
		})
	})

	//  localhost:8080/isChainValid
	r.GET("/isChainValid", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"isValid": blockchain.isChainValid(blockchain.chain),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
