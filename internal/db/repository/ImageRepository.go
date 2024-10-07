package repository

import (
	"vetblock/internal/db"
	"vetblock/internal/db/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ImageRepositoryInterface interface {
	SaveImage(image *model.ImageModel) error
	FindImageByID(id uuid.UUID) (*model.ImageModel, error)
	DeleteImage(id uuid.UUID) (string, error)
	UpdateImage(id uuid.UUID, updatedImage model.ImageModel) error
}

type ImageRepository struct {
	Db *gorm.DB
}

func NewImageRepository() *ImageRepository {
	database := db.NewDb()
	return &ImageRepository{Db: database}
}


func (r *ImageRepository) SaveImage(image *model.ImageModel) error {
	if err := r.Db.Create(image).Error; err != nil {
		return err
	}
	return nil
}

func (r *ImageRepository) FindImageByID(id uuid.UUID) (*model.ImageModel, error) {
	var image model.ImageModel
	if err := r.Db.Where("id = ?", id).First(&image).Error; err != nil {
		return nil, err
	}
	return &image, nil
}

func (r *ImageRepository) DeleteImage(id uuid.UUID) (string, error) {
	var image model.ImageModel
	if err := r.Db.Where("id = ?", id).First(&image).Error; err != nil {
		return "Image not found", err
	}

	if err := r.Db.Delete(&image).Error; err != nil {
		return "Error deleting image", err
	}

	return "Image deleted successfully", nil
}

//update image
func (r *ImageRepository) UpdateImage(id uuid.UUID, updatedImage model.ImageModel) error {
	var image model.ImageModel
	if err := r.Db.Where("id = ?", id).First(&image).Error; err != nil {
		return err
	}

	image.Image = updatedImage.Image

	if err := r.Db.Save(&image).Error; err != nil {
		return err
	}

	return nil
}