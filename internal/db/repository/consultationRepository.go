package repository

import (
	"context"
	"log"
	"vetblock/internal/db"
	"vetblock/internal/db/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConsultationRepository struct {
	db *gorm.DB
}

func GetDB() *gorm.DB {
	return db.NewDb()
}

func NewConsultationRepository(db *gorm.DB) *ConsultationRepository {
	return &ConsultationRepository{db: db}
}

// Método para encontrar consulta por ID
func (repo *ConsultationRepository) FindConsultationByID(ctx context.Context, id uuid.UUID) (*model.Consultation, error) {
	var consultation model.Consultation
	log.Print("Finding consultation by ID REPO")
	result := repo.db.WithContext(ctx).First(&consultation, "id = ?", id)
	if result.Error != nil {
		log.Print("Error finding consultation:", result.Error)
		return nil, result.Error
	}

	return &consultation, nil
}

// Métodos adicionais conforme necessário
func (repo *ConsultationRepository) FindConsultationByAnimalID(ctx context.Context, animalID uuid.UUID) ([]model.Consultation, error) {
	var consultations []model.Consultation
	result := repo.db.WithContext(ctx).Where("animal_id = ?", animalID).Find(&consultations)
	return consultations, result.Error
}

func (repo *ConsultationRepository) FindConsultationByVeterinaryCRVM(ctx context.Context, crvm string) ([]model.Consultation, error) {
	var consultations []model.Consultation
	result := repo.db.WithContext(ctx).Where("crvm = ?", crvm).Find(&consultations)
	return consultations, result.Error
}

func (repo *ConsultationRepository) FindConsultationByDate(ctx context.Context, date string) ([]model.Consultation, error) {
	var consultations []model.Consultation
	result := repo.db.WithContext(ctx).Where("consultation_date = ?", date).Find(&consultations)
	return consultations, result.Error
}

func (repo *ConsultationRepository) FindConsultationByDateRange(ctx context.Context, startDate, endDate string) ([]model.Consultation, error) {
	var consultations []model.Consultation
	result := repo.db.WithContext(ctx).Where("consultation_date BETWEEN ? AND ?", startDate, endDate).Find(&consultations)
	return consultations, result.Error
}

func (repo *ConsultationRepository) FindConsultationByAnimalIDAndDateRange(ctx context.Context, animalID uuid.UUID, startDate, endDate string) ([]model.Consultation, error) {
	var consultations []model.Consultation
	result := repo.db.WithContext(ctx).Where("animal_id = ? AND consultation_date BETWEEN ? AND ?", animalID, startDate, endDate).Find(&consultations)
	return consultations, result.Error
}

// Método para salvar uma consulta
func (repo *ConsultationRepository) SaveConsultation(ctx context.Context, consultation *model.Consultation) error {
	result := repo.db.WithContext(ctx).Save(consultation)
	log.Print("Repository Saving Consultation")
	return result.Error
}

// Método para deletar uma consulta
func (repo *ConsultationRepository) DeleteConsultation(ctx context.Context, id uuid.UUID) error {
	result := repo.db.WithContext(ctx).Delete(&model.Consultation{}, id)
	log.Print("Consultation Deleting Consultation")
	return result.Error
}


//FindConsultationByAnimalIDAndDate
func (repo *ConsultationRepository) FindConsultationByAnimalIDAndDate(ctx context.Context, animalID uuid.UUID, date string) ([]model.Consultation, error) {
	var consultations []model.Consultation
	result := repo.db.WithContext(ctx).Where("animal_id = ? AND consultation_date = ?", animalID, date).Find(&consultations)
	return consultations, result.Error
}

