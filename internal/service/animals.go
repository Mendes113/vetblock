package service

import (
	// "crypto/sha256"
	// "encoding/hex"
	"fmt"
	"log"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"
	"github.com/google/uuid"
)




func  AddAnimalTransaction(animal model.Animal, sender, receiver string) error {
	log.Println("adding animal transaction")


	

	repo := repository.NewAnimalRepository()
	if err := repo.SaveAnimal(&animal); err != nil {
		return err
	}

	return nil
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
