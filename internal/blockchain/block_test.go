package blockchain

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
)

func TestCalculateHash(t *testing.T) {
    block := Block{
        Index:        0,
        Timestamp:    time.Now(),
        Transactions: []Transaction{},
        PreviousHash: "0",
        Hash:         "",
        Nonce:        0,
    }

    // Calcule o hash
    hash := block.CalculateHash()
    
    // Use assert para verificar se o hash não é vazio
    assert.NotEmpty(t, hash, "Expected hash to be non-empty")

    // Verifique o comprimento do hash
    assert.Len(t, hash, 64, "Expected hash length to be 64")
}

func TestMineBlock(t *testing.T) {
	block := Block{
		Index:        0,
		Timestamp:    time.Now(),
		Transactions: []Transaction{},
		PreviousHash: "0",
		Hash:         "",
		Nonce:        0,
	}

	// Mine o bloco com dificuldade 1
	block.MineBlock(1)

	// Use assert para verificar se o hash não é vazio
	assert.NotEmpty(t, block.Hash, "Expected hash to be non-empty")

	// Verifique se o hash atende à dificuldade
	assert.True(t, isHashValid(block.Hash, 1), "Expected hash to meet difficulty")
}

func TestCalculateHashError(t *testing.T){
	block := Block{
		Index:        0,
		Timestamp:    time.Now(),
		Transactions: []Transaction{},
		PreviousHash: "0",
		Hash:         "",
		Nonce:        0,
	}

	// Calcule o hash
	hash := block.CalculateHash()
	stringHash := "12345678901234567890123456789012345678901234567890123456789012345678901234567890"

	// Use assert para verificar se o hash não é vazio
	assert.NotEmpty(t, hash, "Expected hash to be non-empty")
	assert.NotEqual(t, hash, stringHash, "Expected hash to be different")
	// Verifique o comprimento do hash
	assert.Len(t, hash, 64, "Expected hash length to be 64")
}