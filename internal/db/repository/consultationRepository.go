package repository

import (
	"errors"
	"log"
	"vetblock/internal/db"
	"vetblock/internal/db/model"

	"gorm.io/gorm"
)

type ConsultationRepository struct {
	Db *gorm.DB
}

func NewConsultationRepository() *ConsultationRepository {
	return &ConsultationRepository{
		Db: db.NewDb(),
	}
}

func (r *ConsultationRepository) SaveConsultation(consult *model.Consultation) error {
	log.Print("Repository Saving Consultation")
	if err := r.Db.Create(consult).Error; err != nil {
		log.Print("Error saving consultation:", err)
		return err
	}
	log.Print(consult)
	log.Print("Repository Saving Animal")
	return nil
}

func (r *ConsultationRepository) FindConsultationByID(consult model.Consultation) (*model.Consultation, error) {
	if err := r.Db.Where("id = ? AND deleted_at IS NULL", consult.ID).First(&consult).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Print("Consultation not found")
			return nil, nil // ou outro retorno que faça sentido para o seu caso
		}
		log.Print("Error finding consultation:", err)
		return nil, err
	}
	return &consult, nil
}

func (r *ConsultationRepository) FindConsultationByAnimalID(consult model.Consultation) (*model.Consultation, error) {
	if err := r.Db.Where("animal_id = ? AND deleted_at IS NULL", consult.AnimalID).First(&consult).Error; err != nil {
		log.Print("Error finding consultation:", err)
		return nil, err

	}
	return &consult, nil
}

func (r *ConsultationRepository) FindConsultationByVeterinaryCRVM(consult model.Consultation) (*model.Consultation, error) {
	if err := r.Db.Where("veterinary_id = ? AND deleted_at IS NULL", consult.CRVM).First(&consult).Error; err != nil {
		log.Print("Error finding consultation:", err)
		return nil, err

	}
	return &consult, nil
}

func (r *ConsultationRepository) FindConsultationByDate(consult model.Consultation) (*model.Consultation, error) {
	if err := r.Db.Where("date = ? AND deleted_at IS NULL", consult.ConsultationDate).First(&consult).Error; err != nil {
		log.Print("Error finding consultation:", err)
		return nil, err

	}
	return &consult, nil
}

func (r *ConsultationRepository) FindConsultationByDateRange(consult model.Consultation) (*model.Consultation, error) {
	if err := r.Db.Where("date BETWEEN ? AND ? AND deleted_at IS NULL", consult.ConsultationDate, consult.ConsultationDate).First(&consult).Error; err != nil {
		log.Print("Error finding consultation:", err)
		return nil, err

	}
	return &consult, nil
}

func (r *ConsultationRepository) FindConsultationByAnimalIDAndDateRange(consult model.Consultation) (*model.Consultation, error) {
	if err := r.Db.Where("animal_id = ? AND date BETWEEN ? AND ? AND deleted_at IS NULL", consult.AnimalID, consult.ConsultationDate, consult.ConsultationDate).First(&consult).Error; err != nil {
		log.Print("Error finding consultation:", err)
		return nil, err

	}
	return &consult, nil

}

func (r *ConsultationRepository) FindConsultationByVeterinaryIDAndDateRange(consult model.Consultation) (*model.Consultation, error) {
	if err := r.Db.Where("veterinary_id = ? AND date BETWEEN ? AND ? AND deleted_at IS NULL", consult.CRVM, consult.ConsultationDate, consult.ConsultationDate).First(&consult).Error; err != nil {
		log.Print("Error finding consultation:", err)
		return nil, err

	}
	return &consult, nil
}

// is there a consultation at the same time?
func (r *ConsultationRepository) FindConsultationByDateAndVeterinaryID(consult model.Consultation) (*model.Consultation, error) {
	if err := r.Db.Where("date = ? AND veterinary_id = ? AND deleted_at IS NULL", consult.ConsultationDate, consult.CRVM).First(&consult).Error; err != nil {
		log.Print("Error finding consultation:", err)
		return nil, err

	}
	return &consult, nil
}

func (r *ConsultationRepository) DeleteConsultation(consult model.Consultation) (string, error) {
	if err := r.Db.Where("id = ?", consult.ID).First(&consult).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Print("Consultation not found:", err)
			return "Consultation not found", err
		}
		log.Print("Error finding consultation:", err)
		return "Error finding consultation", err
	}

	// Soft delete
	if err := r.Db.Delete(&consult).Error; err != nil {
		log.Print("Error deleting consultation:", err)
		return "Error deleting consultation", err
	}

	log.Print("Consultation deleted successfully")
	return "Consultation deleted successfully", nil
}
