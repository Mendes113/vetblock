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

//criar um novo bloco
func NewBlock(index int, transactions []Transaction, previousHash string) *Block {
    block := &Block{
        Index:        index,
        Timestamp:    time.Now(),
        Transactions: transactions,
        PreviousHash: previousHash,
        Hash:         "",
        Nonce:        0,
    }
    block.Hash = block.CalculateHash()
    return block
}


func ValidateBlock(block Block) bool {
    // Verificar o hash do bloco
    if block.Hash != block.CalculateHash() {
        return false
    }

    // Verificar o hash do bloco anterior
    if len(Blockchain) > 0 {
        lastBlock := Blockchain[len(Blockchain)-1]
        if block.PreviousHash != lastBlock.Hash {
            return false
        }
    }

    // Verificar a prova de trabalho
    difficulty := 4 // Ajuste a dificuldade conforme necessário
    if !isHashValid(block.Hash, difficulty) {
        return false
    }

    // Validar transações
    if !validateTransactions(block.Transactions) {
        return false
    }

    // Verificar o timestamp
    if block.Timestamp.After(time.Now()) {
        return false
    }

    // Verificar índice sequencial
    if len(Blockchain) > 0 {
        lastBlock := Blockchain[len(Blockchain)-1]
        if block.Index != lastBlock.Index+1 {
            return false
        }
    }

    return true
}

// Função para validar transações
func validateTransactions(transactions []Transaction) bool {
    
    for _, transaction := range transactions {
        if !transaction.Validate() {
            return false
        }

        
        
    }



    return true
}