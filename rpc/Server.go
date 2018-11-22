package rpc

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/demochain/core"
)

var blockchain *core.Blockschain

func Run() {
	blockchain = core.NewBlockchain()
	http.HandleFunc("/blockchain/get", blockchainGet)
	http.HandleFunc("/blockchain/write", blockchainWrite)
	http.ListenAndServe("localhost:8888", nil)
}

func blockchainGet(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(blockchain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func blockchainWrite(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGet(w, r)
}
