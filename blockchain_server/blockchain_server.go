package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/uglo1/blockchain/block"
	"github.com/uglo1/blockchain/wallet"
)

/*
 * BlockchainServer
 */

type BlockchainServer struct {
	port uint16
}

var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func (bcs *BlockchainServer) GetBlockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallet := wallet.NewWallet()
		bc := block.NewBlockchain(minersWallet.BlockchainAddress(), bcs.Port())
		cache["blockchain"] = bc
		// REMOVE LATER
		log.Printf("private_key %v", minersWallet.PrivateKeyStr())
		log.Printf("publick_key %v", minersWallet.PublicKeyStr())
		log.Printf("blockchain_address %v", minersWallet.BlockchainAddress())
	}
	return bc
}

func (bcs *BlockchainServer) Getchain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.Getchain)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
