package main

import (
	"encoding/json"
	"fmt"
	"goblockchain/utils"
	"goblockchain/wallet"
	"html/template"
	"io"
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

// Wallet is api to make wallet and return.
func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method.")
	}
}

// CreateTransaction is api to create transaction.
func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t wallet.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Printf("ERROR: %v", err)
			io.WriteString(w, string(utils.JSONStatus("fail")))
			return
		}
		if !t.Validate() {
			log.Println("ERROR: missing field(s)")
			io.WriteString(w, string(utils.JSONStatus("fail")))
			return
		}

		fmt.Println(*t.SenderPublicKey)
		fmt.Println(*t.SenderPrivateKey)
		fmt.Println(*t.SenderBlockchainAddress)
		fmt.Println(*t.RecipientBlockchainAddress)
		fmt.Println(*t.Value)

		// io.WriteString(w, string(utils.JSONStatus("success")))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

// Run is to run wallet server.
func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
