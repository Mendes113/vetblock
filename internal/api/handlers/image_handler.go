package handlers

import (
	
	"vetblock/internal/db/model"
	"vetblock/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ImageHandler struct {
	ImageService *service.ImageService
}

// NewImageHandler cria uma nova instância de ImageHandler
func NewImageHandler(imageService *service.ImageService) *ImageHandler {
	return &ImageHandler{ImageService: imageService}
}

// AddImageHandler lida com a criação de novas imagens
func (h *ImageHandler) AddImageHandler(c *fiber.Ctx) error {
	var image model.ImageModel

	// Decodifica o corpo da requisição JSON para o struct ImageModel
	if err := c.BodyParser(&image); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Chama o serviço para adicionar a imagem
	if err := h.ImageService.AddImage(image); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add image"})
	}

	// Retorna sucesso
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Image added successfully"})
}

// GetImageByIDHandler lida com a recuperação de uma imagem por ID
func (h *ImageHandler) GetImageByIDHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// Converte o parâmetro da URL para UUID
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid image ID"})
	}

	// Chama o serviço para buscar a imagem
	image, err := h.ImageService.GetImageByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Image not found"})
	}

	// Retorna a imagem como JSON
	return c.JSON(image)
}

// DeleteImageHandler lida com a remoção de uma imagem por ID
func (h *ImageHandler) DeleteImageHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// Converte o parâmetro da URL para UUID
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid image ID"})
	}

	// Chama o serviço para deletar a imagem
	message, err := h.ImageService.DeleteImage(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete image"})
	}

	// Retorna sucesso
	return c.JSON(fiber.Map{"message": message})
}

// UpdateImageHandler lida com a atualização de uma imagem por ID
func (h *ImageHandler) UpdateImageHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// Converte o parâmetro da URL para UUID
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid image ID"})
	}

	var updatedImage model.ImageModel
	// Decodifica o corpo da requisição JSON para o struct ImageModel
	if err := c.BodyParser(&updatedImage); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Chama o serviço para atualizar a imagem
	if err := h.ImageService.UpdateImage(id, updatedImage); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update image"})
	}

	// Retorna sucesso
	return c.JSON(fiber.Map{"message": "Image updated successfully"})
}
