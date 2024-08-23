package service

import (
	// "crypto/sha256"
	// "encoding/hex"
	"errors"
	"fmt"
	"log"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"

	"github.com/google/uuid"
)




func AddAnimal(animal model.Animal) error {
    log.Println("adding animal transaction")

    repo := repository.NewAnimalRepository()
    existingAnimal, err := repo.FindByUniqueAttributes(animal)
    if err != nil {
        return err
    }
    if existingAnimal != nil {
        return errors.New("animal já existe")
    }
    
    if err := repo.SaveAnimal(&animal); err != nil {
        return err
    }

    return nil
}

func GetAnimalByID(id uuid.UUID) (*model.Animal, error) {
	repo := repository.NewAnimalRepository()
	animal, err := repo.FindAnimalByID(id)
	if err != nil {
		return nil, err
	}
	return animal, nil
}


// Atualiza um animal na blockchain
func  UpdateAnimal(id uuid.UUID, updatedAnimal model.Animal) error {
	log.Println("Atualizando animal")

	repo := repository.NewAnimalRepository()
	animal, err := repo.FindAnimalByID(id)
	if err != nil {
		return err
	}

	// Atualiza os campos do animal
	animal.Name = updatedAnimal.Name
	animal.Species = updatedAnimal.Species
	animal.Breed = updatedAnimal.Breed
	animal.Age = updatedAnimal.Age
	animal.Description = updatedAnimal.Description
	animal.CPFTutor = updatedAnimal.CPFTutor

	if err := repo.SaveAnimal(animal); err != nil {
		return err
	}
	

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
