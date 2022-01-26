package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

// BlockchainServer is BlockchainServer struct.
type BlockchainServer struct {
	port uint16
}

// NewBlockchainServer is to return new NewBlockchainServer struct.
func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

// Port is to return BlockchainServer's port.
func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

// HelloWorld is write hello world.
func HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, World!")
}

// Run is to run server.
func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
