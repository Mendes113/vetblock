package handlers

import (
	"log"
	"vetblock/internal/db/model"
	"vetblock/internal/service"

	"github.com/gofiber/fiber/v2"
)

type VeterinaryResponse struct {
	Name        string `json:"name" gorm:"not null" validate:"required,min=2,max=100"`
	CRVM        string `json:"crvm" gorm:"not null" validate:"required"`
	PhoneNumber string `json:"phone_number" gorm:"not null" validate:"required"`
	LastName   string `json:"last_name" gorm:"not null" validate:"required"`
	Email       string `json:"email" gorm:"not null" validate:"required,email"`	
}




func AddVeterinaryHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var veterinary VeterinaryResponse
		if err := c.BodyParser(&veterinary); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}
		log.Print(veterinary)
		// Convert the VeterinaryResponse to Veterinary
		veterinaryModel := model.Veterinary{
			CRVM:  veterinary.CRVM,
			Name:  veterinary.Name,
			LastName: veterinary.LastName,	
			Email: veterinary.Email,
			Phone: veterinary.PhoneNumber,
		}
		
		err := service.AddVeterinary(veterinaryModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to add veterinary transaction")
		}

		return c.Status(fiber.StatusCreated).JSON(veterinary)	
	}

}
