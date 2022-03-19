package main

import (
	"fmt"
	"log"
	"time"
)

func NewBlock(nonce int, previoutHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previoutHash
	return b
}
func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	log.Println("test")
	fmt.Println("test2")
}
