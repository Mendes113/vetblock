package repository

import (
	"log"

	"vetblock/internal/db"
	"vetblock/internal/db/model"

	"gorm.io/gorm"
)

type VeterinaryRepository struct {
	Db *gorm.DB
}

func NewVeterinaryRepository() *VeterinaryRepository {
	return &VeterinaryRepository{
		Db: db.NewDb(),
	}
}

func (r *VeterinaryRepository) SaveVeterinary(veterinary model.Veterinary) error {
	log.Println("saving veterinary transaction")
	if err := r.Db.Create(veterinary).Error; err != nil {
		log.Print("Error saving veterinary:", err)
		return err
	}
	log.Print(veterinary)
	log.Print("Repository Saving Veterinary")
	return nil
}

func (r *VeterinaryRepository) FindVeterinaryByCRVM(crvm string) (*model.Veterinary, error) {
	var veterinary model.Veterinary
	if err := r.Db.Where("crvm = ? AND deleted_at IS NULL", crvm).First(&veterinary).Error; err != nil {
		log.Print("Error finding veterinary:", err)
		return nil, err
	}
	return &veterinary, nil
}