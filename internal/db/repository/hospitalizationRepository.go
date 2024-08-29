package repository

import (
	"log"
	"vetblock/internal/db"
	"vetblock/internal/db/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HospitalizationRepository struct {
	Db *gorm.DB
}

func NewHospitalizationRepository() *HospitalizationRepository {
	return &HospitalizationRepository{
		Db: db.NewDb(),
	}
}

// Salva uma hospitalização no banco de dados e retorna um erro se ocorrer
func (r *HospitalizationRepository) SaveHospitalization(hospitalization *model.Hospitalization) error {
	if err := r.Db.Create(hospitalization).Error; err != nil {
		log.Print("Error saving hospitalization:", err)
		return err
	}
	log.Print(hospitalization)
	log.Print("Repository Saving Hospitalization")
	return nil
}

func (r *HospitalizationRepository) FindHospitalizationByID(id string) (*model.Hospitalization, error) {
	var hospitalization model.Hospitalization
	if err := r.Db.Where("id = ? AND deleted_at IS NULL", id).First(&hospitalization).Error; err != nil {
		log.Print("Error finding hospitalization:", err)
		return nil, err
	}
	return &hospitalization, nil
}

// delete hospitalization

func (r *HospitalizationRepository) DeleteHospitalization(id string) (string, error) {
	var hospitalization model.Hospitalization
	if err := r.Db.Where("id = ?", id).First(&hospitalization).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Print("Hospitalization not found:", err)
			return "Hospitalization not found", err
		}
		log.Print("Error finding hospitalization:", err)
		return "Error finding hospitalization", err
	}

	// Soft delete
	if err := r.Db.Delete(&hospitalization).Error; err != nil {
		log.Print("Error deleting hospitalization:", err)
		return "Error deleting hospitalization", err
	}

	log.Print("Hospitalization excluído com sucesso")
	return "Hospitalization deleted successfully", nil
}


func (r *HospitalizationRepository) GetHospitalizationByID(id uuid.UUID) (*model.Hospitalization, error) {
	var hospitalization model.Hospitalization
	if err := r.Db.Where("id = ?", id).First(&hospitalization).Error; err != nil {
		log.Print("Error finding hospitalization:", err)
		return nil, err
	}
	return &hospitalization, nil
}