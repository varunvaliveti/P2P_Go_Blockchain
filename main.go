package main

import (
	"encoding/hex"
	"fmt"
	"github.com/varunvaliveti/P2P_Go_Blockchain"
)

func main() {
	chain := P2P_Go_Blockchain.InitBlockChain()
	chain.AddBlock("First")
	chain.AddBlock("Second")

	for _, block := range chain.blocks {
		// fmt.Println("Prev hash ", hex.EncodeToString(block.PrevHash))
		fmt.Println("data in block ", string(block.Data))
		fmt.Println("Hash: ", hex.EncodeToString(block.Hash))
		fmt.Println()
	}

}
