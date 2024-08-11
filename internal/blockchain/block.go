package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
	 "strconv"
)

// Estrutura do Bloco
type Block struct {
    Index        int           `json:"index"`
    Timestamp    time.Time     `json:"timestamp"`
    Transactions []Transaction `json:"transactions"`
    PreviousHash string        `json:"previous_hash"`
    Hash         string        `json:"hash"`
    Nonce        int           `json:"nonce"`
}

// Função para calcular o hash de um bloco
func (b *Block) CalculateHash() string {
    transactionsJSON, _ := json.Marshal(b.Transactions)
   
    
    record := strconv.Itoa(b.Index) + b.Timestamp.String() + string(transactionsJSON) + b.PreviousHash + strconv.Itoa(b.Nonce)
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}
