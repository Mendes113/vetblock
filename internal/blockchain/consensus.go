package blockchain

import (
    "fmt"
)

// Função para calcular o hash de um bloco com um nonce (Proof of Work)
func (b *Block) MineBlock(difficulty int) {
    nonce := 0
    hash := ""
    for !isHashValid(hash, difficulty) {
        nonce++
        b.Nonce = nonce
        hash = b.CalculateHash()
    }
    b.Hash = hash
}

// Função para verificar se o hash atende à dificuldade (exemplo simples)
func isHashValid(hash string, difficulty int) bool {
    if len(hash) < difficulty {
        return false
    }
    prefix := ""
    for i := 0; i < difficulty; i++ {
        prefix += "0"
    }
    return hash[:difficulty] == prefix
}

// Função para adicionar um bloco à blockchain com Proof of Work
func AddBlockWithConsensus(newBlock Block, difficulty int) error {
    lastBlock := Blockchain[len(Blockchain)-1]
    if newBlock.PreviousHash != lastBlock.Hash {
        return fmt.Errorf("Invalid previous hash")
    }
    newBlock.MineBlock(difficulty)
    Blockchain = append(Blockchain, newBlock)
    return nil
}
