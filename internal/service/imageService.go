package service

import (
	"encoding/base64"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"

	"github.com/google/uuid"
	"fmt"
)

type ImageService struct {
	repo repository.ImageRepositoryInterface
}

// NewImageService creates a new instance of ImageService with a repository
func NewImageService(repo repository.ImageRepositoryInterface) *ImageService {
	return &ImageService{repo: repo}
}

// AddImage adds a new image to the repository
func (s *ImageService) AddImage(image model.ImageModel) error {
	// Converte a imagem (em bytes) para base64 corretamente
	base64Image, err := s.EncodeImageToBase64(image.Image)
	if err != nil {
		return fmt.Errorf("erro ao codificar imagem em base64: %v", err)
	}

	// Converte de volta de base64 para bytes
	image.Image, err = s.DecodeBase64ToBytes(base64Image)
	if err != nil {
		return fmt.Errorf("erro ao decodificar imagem de base64: %v", err)
	}

	return s.repo.SaveImage(&image)
}

// GetImageByID retrieves an image from the repository by its ID
// GetImageByID busca a imagem do repositório e retorna como base64
func (s *ImageService) GetImageByID(id uuid.UUID) (string, error) {
	// Busca a imagem pelo ID no banco de dados
	image, err := s.repo.FindImageByID(id)
	if err != nil {
		return "", err
	}

	// Verifica se a imagem foi encontrada
	if image == nil || len(image.Image) == 0 {
		return "", fmt.Errorf("imagem não encontrada")
	}

	// Converte a imagem em base64
	base64Image := base64.StdEncoding.EncodeToString(image.Image)

	return base64Image, nil
}

func (s *ImageService) DeleteImage(id uuid.UUID) (string, error) {
	return s.repo.DeleteImage(id)
}

func (s *ImageService) UpdateImage(id uuid.UUID, updatedImage model.ImageModel) error {
	return s.repo.UpdateImage(id, updatedImage)
}

// EncodeImageToBase64 converts an image (byte slice) to a base64 encoded string
func (s *ImageService) EncodeImageToBase64(image []byte) (string, error) {
	if len(image) == 0 {
		return "", fmt.Errorf("imagem vazia")
	}

	// Codifica a imagem em base64
	return base64.StdEncoding.EncodeToString(image), nil
}

// DecodeBase64ToBytes converts a base64 string back to a byte slice
func (s *ImageService) DecodeBase64ToBytes(base64Str string) ([]byte, error) {
	// Decodifica a string base64 de volta para bytes
	imageBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar base64: %v", err)
	}
	return imageBytes, nil
}

//decode Bytes to base64
func (s *ImageService) DecodeBytesToBase64(image []byte) (string, error) {
	return base64.StdEncoding.EncodeToString(image), nil
}


//find image and decode to base64
func (s *ImageService) FindImageAndDecodeToBase64(id uuid.UUID) (string, error) {
	image, err := s.GetImageByID(id)
	if err != nil {
		return "", err
	}

	// Decodifica a imagem para base64
	base64Image, err := s.DecodeBytesToBase64([]byte(image))
	if err != nil {
		return "", err
	}

	return base64Image, nil
}