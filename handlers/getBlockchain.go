package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/ssharif6/LoanBlockchain/models"
	"github.com/davecgh/go-spew/spew"
)

func (ctx *HandlerCtx) BlockchainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.Method == "GET" {
		if err := json.NewEncoder(w).Encode(ctx.Blockchain); err != nil {
			http.Error(w, "error unmarshaling blockchain", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		req := &models.Request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, "error decoding body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		prev := ctx.Blockchain.Chain[len(ctx.Blockchain.Chain) - 1]
		nb, err := ctx.Blockchain.AddBlock(prev, req.Amount)
		if err != nil {
			http.Error(w, "error adding block", http.StatusInternalServerError)
			return
		}

		if ctx.Blockchain.IsBlockValid(nb, prev) {
			nbc := append(ctx.Blockchain.Chain, nb)
			ctx.Blockchain.ChooseChain(nbc)
			spew.Dump(ctx.Blockchain)
		}

		if err := json.NewEncoder(w).Encode(nb); err != nil {
			http.Error(w, "error encoding block", http.StatusInternalServerError)
			return
		}

	} else {
		http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
	}
}
