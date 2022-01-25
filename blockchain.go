package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// MiningDifficulty is mining difficulty.
const MiningDifficulty = 3

// Block is block struct.
type Block struct {
	timestamp    int64
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
}

// NewBlock is to return new Block struct.
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		timestamp:    time.Now().UnixNano(),
		transactions: transactions,
	}
}

// Print is print block data.
func (b *Block) Print() {
	fmt.Printf("timestamp             %d\n", b.timestamp)
	fmt.Printf("nonce                 %d\n", b.nonce)
	fmt.Printf("previousHash          %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

// Hash is to return sha256.Sum256 hash.
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

// MarshalJSON is override Block's marshaljson.
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transaction  []*Transaction `json:"transaction"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transaction:  b.transactions,
	})
}

// Blockchain is blockchain struct.
type Blockchain struct {
	transactionPool []*Transaction
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
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
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

// AddTransaction is add transaction to transaction pool
func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

// CopyTransactionPool is to return copy transaction pool.
func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions,
			NewTransaction(t.senderBlockchainAddress,
				t.recipientBlockchainAddress,
				t.value))
	}
	return transactions
}

// ValidProof is validate "000"
func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

// ProofOfWork is proof of work.
func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, MiningDifficulty) {
		nonce++
	}
	return nonce
}

// Transaction is transaction struct.
type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

// NewTransaction is to return new Transaction struct.
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

// Print is print transaction data.
func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("sender_blockchain_address      %s\n", t.senderBlockchainAddress)
	fmt.Printf("recipient_blockchain_address   %s\n", t.recipientBlockchainAddress)
	fmt.Printf("value                          %.2f\n", t.value)
}

// MarshalJSON is override Transaction's marshaljson.
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockChain := NewBlockchain()

	blockChain.AddTransaction("A", "B", 1.2)
	blockChain.AddTransaction("A", "B", 1.5)
	previousHash := blockChain.LastBlock().Hash()
	nonce := blockChain.ProofOfWork()
	blockChain.CreateBlock(nonce, previousHash)

	blockChain.AddTransaction("C", "X", 2.2)
	nonce = blockChain.ProofOfWork()
	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(nonce, previousHash)
	blockChain.Print()
}
