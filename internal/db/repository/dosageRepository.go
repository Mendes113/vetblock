package repository

import (
	"context"
	"vetblock/internal/db/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DosageRepository interface {
	Create(ctx context.Context, dosage *model.Dosage) error
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

func (r *dosageRepository) Create(ctx context.Context, dosage *model.Dosage) error {
	return r.db.WithContext(ctx).Create(dosage).Error
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
