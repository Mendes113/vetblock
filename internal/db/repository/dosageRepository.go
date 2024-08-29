package repository

import (
	"context"
	"log"
	"vetblock/internal/db/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DosageRepository interface {
	Create(ctx context.Context, dosage *model.Dosage, medicationId uuid.UUID, quantity int) error
	Update(ctx context.Context, dosage *model.Dosage) error
	Delete(ctx context.Context, dosageID uuid.UUID) error
	FindByID(ctx context.Context, dosageID uuid.UUID) (*model.Dosage, error)
	FindByAnimalID(ctx context.Context, animalID uuid.UUID) ([]model.Dosage, error)
}

type dosageRepository struct {
	db *gorm.DB
}

func NewDosageRepository(db *gorm.DB) DosageRepository {
	return &dosageRepository{db: db}
}

func (r *dosageRepository) Create(ctx context.Context, dosage *model.Dosage, medicationId uuid.UUID,quantity int) error {
    // Inicia uma transação
    tx := r.db.WithContext(ctx).Begin()

    // Tenta criar o novo item na tabela de dosagens
    if err := tx.Create(dosage).Error; err != nil {
        tx.Rollback() // Reverte a transação em caso de erro
        return err
    }
	log.Print("Repository Saving Dosage")
    // Atualiza a quantidade de medicamentos disponíveis
	if err := tx.Model(&model.Medication{}).Where("id = ?", medicationId).Update("quantity", gorm.Expr("quantity - ?", quantity)).Error; err != nil {
		tx.Rollback() // Reverte a transação em caso de erro
		return err
	}	
	log.Print("Repository Updating Medication Quantity")


    // Confirma a transação se tudo ocorreu sem erros
    return tx.Commit().Error
}

func (r *dosageRepository) Update(ctx context.Context, dosage *model.Dosage) error {
	return r.db.WithContext(ctx).Save(dosage).Error
}

func (r *dosageRepository) Delete(ctx context.Context, dosageID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.Dosage{}, dosageID).Error
}

func (r *dosageRepository) FindByID(ctx context.Context, dosageID uuid.UUID) (*model.Dosage, error) {
	var dosage model.Dosage
	err := r.db.WithContext(ctx).First(&dosage, dosageID).Error
	if err != nil {
		return nil, err
	}
	return &dosage, nil
}

func (r *dosageRepository) FindByAnimalID(ctx context.Context, animalID uuid.UUID) ([]model.Dosage, error) {
	var dosages []model.Dosage
	err := r.db.WithContext(ctx).Where("animal_id = ?", animalID).Find(&dosages).Error
	if err != nil {
		return nil, err
	}
	return dosages, nil
}
