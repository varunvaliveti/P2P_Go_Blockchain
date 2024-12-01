package main

import (
	"fmt"
	"strconv"

	"github.com/varunvaliveti/P2P_Go_Blockchain/Blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.Blocks {

		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in previous block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("Proof of Work: %s\n\n", strconv.FormatBool(pow.Validate()))

	}

}
