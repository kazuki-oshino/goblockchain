package main

import (
	"fmt"
	"log"
	"strings"
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

// Blockchain is blockchain struct.
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// NewBlockchain is to return new Blockchain struct.
func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

// CreateBlock is to return new Block struct.
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

// Print is print blockchain data.
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockChain := NewBlockchain()
	blockChain.CreateBlock(2, "hash 1")
	blockChain.CreateBlock(7, "hash 2")
	blockChain.Print()
}
