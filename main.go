package main

import (
	"net/http"
	"os"
	"log"
	"github.com/ssharif6/LoanBlockchain/handlers"
	"github.com/ssharif6/LoanBlockchain/models"
	"time"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	setupServer()
}
func setupServer() {
	genesisBlock := &models.Block{
		Timestamp: time.Now().String(),
		Index: 0,
		Hash: "",
		PrevHash: "",
	}
	spew.Dump(genesisBlock)
	handlerCtx := &handlers.HandlerCtx{
		Blockchain: &models.Blockchain{
			Chain: []*models.Block{genesisBlock},
		},
	}

	mux := http.NewServeMux()
	httpAddr := os.Getenv("ADDR")
	mux.HandleFunc("/v1/blockchain", handlerCtx.BlockchainHandler)
	log.Printf("Listening to %s", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, mux))
}
