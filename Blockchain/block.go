package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	Blocks []*Block
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
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)

	chain.Blocks = append(chain.Blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("First", nil)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{Genesis()},
	}
}
