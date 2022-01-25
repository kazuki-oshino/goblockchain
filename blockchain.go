package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// Block is block struct.
type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []string
}

// NewBlock is to return new Block struct.
func NewBlock(nonce int, previousHash [32]byte) *Block {
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
	fmt.Printf("previousHash          %x\n", b.previousHash)
	fmt.Printf("transactions          %s\n", b.transactions)
}

// Hash is to return sha256.Sum256 hash.
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

// MarshalJSON is override Block's marshaljson.
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PreviousHash [32]byte `json:"previous_hash"`
		Transaction  []string `json:"transaction"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transaction:  b.transactions,
	})
}

// Blockchain is blockchain struct.
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// NewBlockchain is to return new Blockchain struct.
func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

// CreateBlock is to return new Block struct.
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

// LastBlock is find last block at chain.
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
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

	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(2, previousHash)

	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(7, previousHash)
	blockChain.Print()
}
