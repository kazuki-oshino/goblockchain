package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

const tempDir = "templates"

// WalletServer is WalletServer struct.
type WalletServer struct {
	port    uint16
	gateway string
}

// NewWalletServer is to return new wallet server struct.
func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

// Port is return to Wallet port.
func (ws *WalletServer) Port() uint16 {
	return ws.port
}

// Gateway is return to Wallet Gateway.
func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

// Index is index.
func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(tempDir, "index.html"))
		t.Execute(w, "")
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

// Run is to run wallet server.
func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
