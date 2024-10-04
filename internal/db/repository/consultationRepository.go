package repository

import (
	"context"
	"log"
	"vetblock/internal/db"
	"vetblock/internal/db/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Interface ConsultationRepository define os métodos para manipulação das consultas
type ConsultationRepository interface {
	FindConsultationByID(ctx context.Context, id uuid.UUID) (*model.Consultation, error)
	SaveConsultation(ctx context.Context, consultation *model.Consultation) error
	DeleteConsultation(ctx context.Context, id uuid.UUID) error
	FindConsultationByAnimalID(ctx context.Context, animalID uuid.UUID) ([]model.Consultation, error)
	FindConsultationByVeterinaryCRVM(ctx context.Context, crvm string) ([]model.Consultation, error)
	FindConsultationByDate(ctx context.Context, date string) ([]model.Consultation, error)
	FindConsultationByDateRange(ctx context.Context, startDate, endDate string) ([]model.Consultation, error)
	FindConsultationByAnimalIDAndDateRange(ctx context.Context, animalID uuid.UUID, startDate, endDate string) ([]model.Consultation, error)
	FindConsultationByAnimalIDAndDate(ctx context.Context, animalID uuid.UUID, date string) ([]model.Consultation, error)
}

// Estrutura ConsultationRepositoryImpl que implementa a interface ConsultationRepository
type ConsultationRepositoryImpl struct {
	db *gorm.DB
}

// Função para obter a instância do banco de dados
func GetDB() *gorm.DB {
	return db.NewDb()
}

// Função para criar uma nova instância do ConsultationRepositoryImpl
func NewConsultationRepository(db *gorm.DB) ConsultationRepository {
	return &ConsultationRepositoryImpl{db: db}
}

// Método para encontrar consulta por ID
func (repo *ConsultationRepositoryImpl) FindConsultationByID(ctx context.Context, id uuid.UUID) (*model.Consultation, error) {
	var consultation model.Consultation
	log.Print("Finding consultation by ID REPO")
	result := repo.db.WithContext(ctx).First(&consultation, "id = ?", id)
	if result.Error != nil {
		log.Print("Error finding consultation:", result.Error)
		return nil, result.Error
	}

	return &consultation, nil
}

// Método para encontrar consulta por animalID
func (repo *ConsultationRepositoryImpl) FindConsultationByAnimalID(ctx context.Context, animalID uuid.UUID) ([]model.Consultation, error) {
	var consultations []model.Consultation
	result := repo.db.WithContext(ctx).Where("animal_id = ?", animalID).Find(&consultations)
	return consultations, result.Error
}

// Método para encontrar consulta por CRVM
func (repo *ConsultationRepositoryImpl) FindConsultationByVeterinaryCRVM(ctx context.Context, crvm string) ([]model.Consultation, error) {
	var consultations []model.Consultation
	result := repo.db.WithContext(ctx).Where("crvm = ?", crvm).Find(&consultations)
	return consultations, result.Error
}

// Método para encontrar consulta por data
func (repo *ConsultationRepositoryImpl) FindConsultationByDate(ctx context.Context, date string) ([]model.Consultation, error) {
	var consultations []model.Consultation
	result := repo.db.WithContext(ctx).Where("consultation_date = ?", date).Find(&consultations)
	return consultations, result.Error
}

// Método para encontrar consultas em um intervalo de datas
func (repo *ConsultationRepositoryImpl) FindConsultationByDateRange(ctx context.Context, startDate, endDate string) ([]model.Consultation, error) {
	var consultations []model.Consultation
	result := repo.db.WithContext(ctx).Where("consultation_date BETWEEN ? AND ?", startDate, endDate).Find(&consultations)
	return consultations, result.Error
}

// Método para encontrar consultas por animal e intervalo de datas
func (repo *ConsultationRepositoryImpl) FindConsultationByAnimalIDAndDateRange(ctx context.Context, animalID uuid.UUID, startDate, endDate string) ([]model.Consultation, error) {
	var consultations []model.Consultation
	result := repo.db.WithContext(ctx).Where("animal_id = ? AND consultation_date BETWEEN ? AND ?", animalID, startDate, endDate).Find(&consultations)
	return consultations, result.Error
}

// Método para salvar uma consulta
func (repo *ConsultationRepositoryImpl) SaveConsultation(ctx context.Context, consultation *model.Consultation) error {
	result := repo.db.WithContext(ctx).Save(consultation)
	log.Print("Repository Saving Consultation")
	return result.Error
}

// Método para deletar uma consulta
func (repo *ConsultationRepositoryImpl) DeleteConsultation(ctx context.Context, id uuid.UUID) error {
	result := repo.db.WithContext(ctx).Delete(&model.Consultation{}, id)
	log.Print("Repository Deleting Consultation")
	return result.Error
}

// Método para encontrar consultas por animalID e data
func (repo *ConsultationRepositoryImpl) FindConsultationByAnimalIDAndDate(ctx context.Context, animalID uuid.UUID, date string) ([]model.Consultation, error) {
	var consultations []model.Consultation
	result := repo.db.WithContext(ctx).Where("animal_id = ? AND consultation_date = ?", animalID, date).Find(&consultations)
	return consultations, result.Error
}
