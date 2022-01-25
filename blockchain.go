package main

import (
	"fmt"
	"log"
	"time"
)

// Block is block struct.
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

// NewBlock is to return new Block struct.
func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		timestamp:    time.Now().UnixNano(),
	}
}

// Print is print block data.
func (b *Block) Print() {
	fmt.Printf("timestamp             %d\n", b.timestamp)
	fmt.Printf("nonce                 %d\n", b.nonce)
	fmt.Printf("previousHash          %s\n", b.previousHash)
	fmt.Printf("transactions          %s\n", b.transactions)
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	b := NewBlock(34, "string hash")
	b.Print()
}
