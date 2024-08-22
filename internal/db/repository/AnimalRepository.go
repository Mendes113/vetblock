package repository

import (
	"errors"
	"log"
	"vetblock/internal/db"
	"vetblock/internal/db/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnimalRepository struct {
	Db *gorm.DB
}

func NewAnimalRepository() *AnimalRepository {
	database := db.NewDb()
	return &AnimalRepository{Db: database}
}

// Salva um animal no banco de dados e retorna um erro se ocorrer
func (r *AnimalRepository) SaveAnimal(animal *model.Animal) error {
	if err := r.Db.Create(animal).Error; err != nil {
		log.Print("Error saving animal:", err)
		return err
	}
	log.Print(animal)
	log.Print("Repository Saving Animal")
	return nil
}

func (r *AnimalRepository) FindAnimalByID(id uuid.UUID) (*model.Animal, error) {
	var animal model.Animal
	if err := r.Db.Where("id = ? AND deleted_at IS NULL", id).First(&animal).Error; err != nil {
		log.Print("Error finding animal:", err)
		return nil, err
	}
	return &animal, nil
}

// delete animal
func (r *AnimalRepository) DeleteAnimal(id uuid.UUID) (string,error) {
	var animal model.Animal
	if err := r.Db.Where("id = ?", id).First(&animal).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Print("Animal not found:", err)
			return "Animal not found", err
		}
		log.Print("Error finding animal:", err)
		return "Error finding animal", err
	}

	// Soft delete
	if err := r.Db.Delete(&animal).Error; err != nil {
		log.Print("Error deleting animal:", err)
		return "Error deleting animal", err
	}

	log.Print("Animal excluído com sucesso")
	return "Animal deleted successfully", nil
}