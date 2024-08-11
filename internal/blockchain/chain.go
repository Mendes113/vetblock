package blockchain

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)




func CreateNewBlock(transactions []Transaction) Block {
    lastBlock := Blockchain[len(Blockchain)-1]
    newBlock := Block{
        Index:        lastBlock.Index + 1,
        Timestamp:    time.Now(),
        Transactions: transactions,
        PreviousHash: lastBlock.Hash,
        Hash:         "",
        Nonce:        0,
    }
    newBlock.Hash = newBlock.CalculateHash()
    return newBlock
}

// Blockchain é uma cadeia de blocos
var Blockchain []Block


// Função para criar um novo bloco com transações

// Função para inicializar a blockchain com o bloco gênese
func InitializeBlockchain() {
    genesisBlock := Block{
        Index:        0,
        Timestamp:    time.Now(),
        Transactions: []Transaction{},
        PreviousHash: "0",
        Hash:         "",
        Nonce:        0,
    }
    genesisBlock.Hash = genesisBlock.CalculateHash()
    Blockchain = append(Blockchain, genesisBlock)
}
// Função para adicionar um novo bloco à blockchain
func AddBlock(newBlock Block) error {
    lastBlock := Blockchain[len(Blockchain)-1]
    if newBlock.PreviousHash != lastBlock.Hash {
        return errors.New("Invalid previous hash")
    }
    newBlock.Hash = newBlock.CalculateHash()
    Blockchain = append(Blockchain, newBlock)
    return nil
}

// Função para validar a blockchain
func ValidateBlockchain() bool {
    for i := 1; i < len(Blockchain); i++ {
        currentBlock := Blockchain[i]
        previousBlock := Blockchain[i-1]

        // Verificar o hash do bloco atual
        if currentBlock.Hash != currentBlock.CalculateHash() {
            return false
        }

        // Verificar o hash do bloco anterior
        if currentBlock.PreviousHash != previousBlock.Hash {
            return false
        }
    }
    return true
}

// Função para obter a cadeia como JSON
func GetBlockchainAsJSON() (string, error) {
    blockchainJSON, err := json.MarshalIndent(Blockchain, "", "    ")
    if err != nil {
        return "", err
    }
    return string(blockchainJSON), nil
}

func GetBlockByIndex(index string) (Block, error) {
	for _, block := range Blockchain {
		if strconv.Itoa(block.Index) == index {
			return block, nil
		}
	}
	return Block{}, errors.New("Block not found")
}


