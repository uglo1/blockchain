package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	timestamp    int64
	nonce        int
	previousHash string
	transactions []string
}

func NewBlock(nonce int, previoutHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previoutHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp           %d\n", b.timestamp)
	fmt.Printf("nonce               %d\n", b.nonce)
	fmt.Printf("previous_hash       %s\n", b.previousHash)
	fmt.Printf("transactions        %s\n", b.transactions)
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}
