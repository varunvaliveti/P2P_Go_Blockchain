package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

//For proof of work, we nee to take the data from the current block
//Make a counter that starts -> 0
// Create hash of the data + the counter
// Verify that the created hash meets certain requirements
// if the hash meets those requirements, we say the hash signs the block
// otherwise we go back and create a new hash
// repeat the step above until we meet the requirements

// The Requirements:
// First few bytes must contains 0s
// The # of 0's before the has is called as the difficulty

const Difficulty = 18 // we make it static, but in real blockchain you have some sort of algorithm that increases difficulty overtime bc #1 you have more miners overtime
//  #2 there is more computing power so you want the difficulty to be relatively the same.

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {

	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow

}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {

		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		intHash.SetBytes(hash[:])

		// now that we hashed it we need to comapre this against our target

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}

	}

	fmt.Println()

	return nonce, hash[:]
}

// this prepeares the data b4 we ues 256 to hash it
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.HashTransactions(),
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)

	intHash.SetBytes(hash[:])

	if intHash.Cmp(pow.Target) == -1 {
		return true
	}

	return false
}

func ToHex(num int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}
