package handlers

import (
	"log"
	"vetblock/internal/db/model"
	"vetblock/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


type AnimalResponse struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name" gorm:"not null" validate:"required,min=2,max=100"`
	Species     string          `json:"species" gorm:"not null" validate:"required"`
	Breed       string          `json:"breed" gorm:"not null" validate:"required"`
	Age         int             `json:"age" validate:"gte=0"`
	Description string          `json:"description"`
	CPFTutor    string          `gorm:"type:char(11);not null" json:"cpf_tutor" validate:"required,len=11"`
}



func AddAnimalTransactionHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var animal AnimalResponse
		animal.ID = uuid.New()
		if err := c.BodyParser(&animal); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		
		// Convert the AnimalResponse to Animal
		animalModel := model.Animal{
			ID:          animal.ID,
			Name:        animal.Name,
			Species:     animal.Species,
			Breed:       animal.Breed,
			Age:         animal.Age,
			Description: animal.Description,
			CPFTutor:    animal.CPFTutor,
		}

		sender := "System"
		receiver := "User"
	

		err := service.AddAnimalTransaction(animalModel, sender, receiver)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to add animal transaction")
		}

		return c.Status(fiber.StatusCreated).JSON(animal)
	}
}

// func GetAnimalByIDHandler() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		id, err := uuid.Parse(c.Params("id"))
// 		if err != nil {
// 			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
// 		}

// 		animal, err := service.GetAnimalByID(id)
// 		if err != nil {
// 			return c.Status(fiber.StatusNotFound).SendString(err.Error())
// 		}

// 		return c.JSON(animal)
// 	}
// }

func UpdateAnimalHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
		}

		var animal model.Animal
		if err := c.BodyParser(&animal); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}
		
		if err := service.UpdateAnimal(id, animal); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to update animal")
		}

		return c.Status(fiber.StatusOK).JSON(animal)
	}
}

func DeleteAnimalHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("Deleting animal")

		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
		}

		msg, err := service.DeleteAnimal(id)
		
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})			
		}
		

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": msg,
		})
	}
}
