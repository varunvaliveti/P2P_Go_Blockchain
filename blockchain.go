package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	hash        string
	timeStamp   time.Time
	proofOfWork int
	data        map[string]interface{}
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func (b Block) calculateHash() {

}
