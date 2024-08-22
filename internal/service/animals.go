package service

import (
	// "crypto/sha256"
	// "encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
	"vetblock/internal/blockchain"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"
	"github.com/google/uuid"
)




func  AddAnimalTransaction(animal model.Animal, sender, receiver string, amount float64) error {
	log.Println("adding animal transaction")

	// Converte o animal para JSON
	animalJSON, err := json.Marshal(animal)
	if err != nil {
		return err // Retorna o erro se a conversão falhar
	}

	// Cria uma nova transação para o animal com o JSON
	transaction := blockchain.Transaction{
		Sender:    sender,
		Receiver:  receiver,
		Timestamp: time.Now(),
		Data:      string(animalJSON), // Armazena o JSON como uma string
	}

	// Cria um novo bloco com a nova transação
	newBlock := blockchain.Block{
		Index:        len(blockchain.Blockchain) + 1,
		Timestamp:    time.Now(),
		Transactions: []blockchain.Transaction{transaction},
		PreviousHash: blockchain.Blockchain[len(blockchain.Blockchain)-1].Hash,
	}

	// Minerar o bloco e adicioná-lo à blockchain
	difficulty := 2
	newBlock.MineBlock(difficulty)
	blockchain.Blockchain = append(blockchain.Blockchain, newBlock)

	repo := repository.NewAnimalRepository()
	if err := repo.SaveAnimal(&animal); err != nil {
		return err
	}

	return nil
}

// Função para buscar um animal por ID na blockchain
func  GetAnimalByID(id uuid.UUID) (*model.Animal, error) {
	// Itera sobre cada bloco na blockchain
	for _, block := range blockchain.Blockchain {
		// Itera sobre cada transação no bloco
		for _, tx := range block.Transactions {
			// Deserializa o campo Data para verificar se contém um animal
			var animal model.Animal
			err := json.Unmarshal([]byte(tx.Data), &animal)
			if err != nil {
				return nil, err // Retorna o erro se a deserialização falhar
			}

			// Verifica se o ID do animal corresponde ao ID pesquisado
			if animal.ID == id {
				return &animal, nil
			}
		}
	}

	return nil, errors.New("animal not found") // Retorna erro se o animal não for encontrado
}

// Atualiza um animal na blockchain
func  UpdateAnimal(id uuid.UUID, updatedAnimal model.Animal) error {
	// Adicione a lógica para atualizar um animal existente
	// Isso pode envolver encontrar o bloco ou transação correspondente e atualizar as informações
	fmt.Printf("Atualizando animal %s com %v\n", id, updatedAnimal)

	return nil
}

// Exclui um animal da blockchain
func  DeleteAnimal(id uuid.UUID) (string, error) {
	repo := repository.NewAnimalRepository()
	fmt.Printf("Excluindo animal %s\n", id)
	if msg, err := repo.DeleteAnimal(id); err != nil {
		return msg, err
	}
	
	return "Animal excluído com sucesso" , nil
}

// // Calcula o hash de uma string
// func calculateHash(data string) string {
//     hash := sha256.New()
//     hash.Write([]byte(data))
//     return hex.EncodeToString(hash.Sum(nil))
// }
