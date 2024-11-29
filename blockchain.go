package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	dataForHash := [][]byte{b.Data, b.PrevHash}
	info := bytes.Join(dataForHash, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]

}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{

		Data:     []byte(data),
		PrevHash: prevHash,
	}

	block.DeriveHash()

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)

	chain.blocks = append(chain.blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("First", nil)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{
		blocks: []*Block{Genesis()},
	}
}

func main() {
	chain := InitBlockChain()
	chain.AddBlock("First")
	chain.AddBlock("Second")

	for _, block := range chain.blocks {
		// fmt.Println("Prev hash ", hex.EncodeToString(block.PrevHash))
		fmt.Println("data in block ", string(block.Data))
		fmt.Println("Hash: ", hex.EncodeToString(block.Hash))
		fmt.Println()
	}

}
