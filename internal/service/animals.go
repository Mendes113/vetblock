package service

import (
	"errors"
	"fmt"
	"log"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"

	"github.com/google/uuid"
)

type AnimalService struct {
	repo repository.AnimalRepositoryInterface
}

// NewAnimalService creates a new instance of AnimalService with a repository
func NewAnimalService(repo repository.AnimalRepositoryInterface) *AnimalService {
	if repo == nil {
		log.Fatal("AnimalService requires a non-nil repository")
	}
	return &AnimalService{repo: repo}
}

// ValidateAnimal checks if the animal has valid data (e.g., a name and species).
func ValidateAnimal(animal model.Animal) error {
	if animal.Name == "" {
		return errors.New("animal needs a name")
	}
	if animal.Species == "" {
		return errors.New("animal needs a species")
	}
	return nil
}

// ValidateAnimalExists checks if the animal already exists in the repository.
func (s *AnimalService) ValidateAnimalExists(animal model.Animal) error {
	existingAnimal, err := s.repo.FindByUniqueAttributes(animal)
	if err != nil {
		return fmt.Errorf("error checking for existing animal: %v", err)
	}
	if existingAnimal != nil {
		return errors.New("animal already exists")
	}
	return nil
}

// AddAnimal adds a new animal to the repository after validation.
func (s *AnimalService) AddAnimal(animal model.Animal) error {
	log.Println("Starting add animal transaction")
	
	// Validate the animal's fields
	if err := ValidateAnimal(animal); err != nil {
		return err
	}

	// Validate if the animal already exists
	if err := s.ValidateAnimalExists(animal); err != nil {
		return err
	}

	// Save the new animal
	if err := s.repo.SaveAnimal(&animal); err != nil {
		return fmt.Errorf("error saving animal: %v", err)
	}

	log.Println("Animal added successfully")
	return nil
}

// GetAnimalByID retrieves an animal by its ID from the repository.
func (s *AnimalService) GetAnimalByID(id uuid.UUID) (*model.Animal, error) {
	animal, err := s.repo.FindAnimalByID(id)
	if err != nil {
		if err.Error() == "animal not found" { // Handle "not found" case specifically
			return nil, fmt.Errorf("animal with ID %s not found", id)
		}
		return nil, err // Propagate other errors
	}
	return animal, nil
}

// UpdateAnimal updates an existing animal's details in the repository.
func (s *AnimalService) UpdateAnimal(id uuid.UUID, updatedAnimal model.Animal) error {
	log.Println("Updating animal")

	animal, err := s.repo.FindAnimalByID(id)
	if err != nil {
		return fmt.Errorf("error finding animal to update: %v", err)
	}

	// Update the animal's fields
	animal.Name = updatedAnimal.Name
	animal.Species = updatedAnimal.Species
	animal.Breed = updatedAnimal.Breed
	animal.Age = updatedAnimal.Age
	animal.Description = updatedAnimal.Description
	animal.CPFTutor = updatedAnimal.CPFTutor

	// Save the updated animal
	if err := s.repo.SaveAnimal(animal); err != nil {
		return fmt.Errorf("error updating animal: %v", err)
	}

	log.Println("Animal updated successfully")
	return nil
}

// DeleteAnimal removes an animal by its ID from the repository.
func (s *AnimalService) DeleteAnimal(id uuid.UUID) (string, error) {
	log.Printf("Deleting animal with ID: %s\n", id)

	msg, err := s.repo.DeleteAnimal(id)
	if err != nil {
		return msg, fmt.Errorf("error deleting animal: %v", err)
	}

	log.Println("Animal deleted successfully")
	return "Animal deleted successfully", nil
}

// GetAllAnimals retrieves all animals from the repository.
func (s *AnimalService) GetAllAnimals() ([]model.Animal, error) {
	animals, err := s.repo.FindAllAnimals()
	if err != nil {
		log.Printf("Error retrieving all animals: %v\n", err)
		return nil, fmt.Errorf("error retrieving animals: %v", err)
	}
	return animals, nil
}
