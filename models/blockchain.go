package models

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index int
	Timestamp string
	Amount string
	Hash string
	PrevHash string
}

type Blockchain struct {
	Chain []*Block
}

// Post request struct
type Request struct {
	Amount string `json:"amount"`
}


func (bc *Blockchain) hash(block *Block) string {
	record := string(block.Index) + block.Timestamp + block.Amount + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

func (bc *Blockchain) AddBlock(prev *Block, amount string) (*Block, error) {
	nb := &Block{}
	nb.Index = prev.Index + 1
	nb.Timestamp = time.Now().String()
	nb.Amount = amount
	nb.PrevHash = prev.Hash
	nb.Hash = bc.hash(nb)

	return nb, nil
}

func (bc *Blockchain) IsBlockValid(nb *Block, prev *Block) bool {
	if nb.Index != prev.Index + 1 {
		return false
	}

	if nb.PrevHash != prev.Hash {
		return false
	}

	if bc.hash(nb) != nb.Hash {
		return false
	}

	return true
}

// Choose longest chain in case of race
func (bc *Blockchain) ChooseChain(nb []*Block) {
	if len(nb) > len(bc.Chain) {
		bc.Chain = nb
	}
}

