package service

import (
	"errors"
	"fmt"
	"log"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"

	"github.com/google/uuid"
)

var animalRepo repository.AnimalRepositoryInterface

// SetAnimalRepository permite injetar um repositório customizado (útil para testes)
func SetAnimalRepository(repo repository.AnimalRepositoryInterface) {
	animalRepo = repo
}

//animal needs at least a name and a species
func ValidateAnimal(animal model.Animal) error {
	if animal.Name == "" {
		return errors.New("animal needs a name")
	}
	if animal.Species == "" {
		return errors.New("animal needs a species")
	}
	return nil
}

//validate if animal already exists
func ValidateAnimalExists(animal model.Animal) error {
	existingAnimal, err := animalRepo.FindByUniqueAttributes(animal)
	if err != nil {
		return err
	}
	if existingAnimal != nil {
		return errors.New("animal already exists")
	}
	return nil
}

func AddAnimal(animal model.Animal) error {
	log.Println("adding animal transaction")

	// Validar se o animal já existe
	if err := ValidateAnimalExists(animal); err != nil {
		return err
	}

	// Salvar o novo animal
	if err := animalRepo.SaveAnimal(&animal); err != nil {
		return err
	}

	return nil
}

func GetAnimalByID(id uuid.UUID) (*model.Animal, error) {
	animal, err := animalRepo.FindAnimalByID(id)
	if err != nil {
		return nil, err
	}
	return animal, nil
}

// Atualiza um animal na blockchain
func UpdateAnimal(id uuid.UUID, updatedAnimal model.Animal) error {
	log.Println("Atualizando animal")

	animal, err := animalRepo.FindAnimalByID(id)
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

	if err := animalRepo.SaveAnimal(animal); err != nil {
		return err
	}

	return nil
}

// Exclui um animal da blockchain
func DeleteAnimal(id uuid.UUID) (string, error) {
	fmt.Printf("Excluindo animal %s\n", id)
	msg, err := animalRepo.DeleteAnimal(id)
	if err != nil {
		return msg, err
	}
	return "Animal excluído com sucesso", nil
}

func GetAllAnimals() ([]model.Animal, error) {
	animals, err := animalRepo.FindAllAnimals()
	if err != nil {
		return nil, err
	}
	return animals, nil
}
