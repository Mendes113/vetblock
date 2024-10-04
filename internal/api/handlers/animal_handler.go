package handlers

import (
	"context"
	"errors"
	"log"
	"time"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"
	"vetblock/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var animal_service = service.NewAnimalService(repository.NewAnimalRepository())

type AnimalResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" gorm:"not null" validate:"required,min=2,max=100"`
	Species     string    `json:"species" gorm:"not null" validate:"required"`
	Breed       string    `json:"breed" gorm:"not null" validate:"required"`
	Weight      float64   `json:"weight" validate:"gte=0"`
	Age         int       `json:"age" validate:"gte=0"`
	Description string    `json:"description"`
	CPFTutor    string    `gorm:"type:char(11);not null" json:"cpf_tutor" validate:"required,len=11"`
}

func AddAnimalHandler() fiber.Handler {
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
			Weight:      animal.Weight,
			Age:         animal.Age,
			Description: animal.Description,
			CPFTutor:    animal.CPFTutor,
		}
		if err := service.ValidateAnimal(animalModel); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		err := animal_service.AddAnimal(animalModel)
		if err != nil {
			//show the error
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
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

		if err := animal_service.UpdateAnimal(id, animal); err != nil {
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

		msg, err := animal_service.DeleteAnimal(id)

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

//add dosage used in animal

type DosageResponse struct {
	ID                uuid.UUID `json:"id"`
	AnimalID          uuid.UUID `json:"animal_id" validate:"required"`
	MedicationID      uuid.UUID `json:"medication_id" validate:"required"`
	StartDate         time.Time `json:"start_date" validate:"required"`
	EndDate           time.Time `json:"end_date" validate:"required"`
	Quantity          int       `json:"quantity" validate:"gte=0"`
	Dosage            string    `json:"dosage" validate:"required"`
	ConsultationID    uuid.UUID `json:"consultation_id"`
	HospitalizationID uuid.UUID `json:"hospitalization_id"`
}

func AddDosageHandler(dosageService *service.DosageService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dosage DosageResponse
		if err := c.BodyParser(&dosage); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		// Create a new UUID for the dosage
		dosageID := uuid.New()

		// Convert the DosageResponse to Dosage
		dosageModel := model.Dosage{
			ID:                dosageID,
			AnimalID:          dosage.AnimalID,
			MedicationID:      dosage.MedicationID,
			StartDate:         dosage.StartDate,
			EndDate:           dosage.EndDate,
			Quantity:          dosage.Quantity,
			Dosage:            dosage.Dosage,
			ConsultationID:    &dosage.ConsultationID,
			HospitalizationID: &dosage.HospitalizationID,
		}

		// Verifica se o animal existe
		animal, err := animal_service.GetAnimalByID(dosage.AnimalID)
		if err != nil {
			return err
		}

		if animal == nil {
			return errors.New("animal não encontrado")
		}

		// Call the service to add the dosage
		err = dosageService.AddDosage(context.Background(), &dosageModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to add dosage transaction")
		}

		// Return the created dosage
		dosage.ID = dosageID
		return c.Status(fiber.StatusCreated).JSON(dosage)
	}
}

func GetAllAnimalsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if context is nil (unlikely but possible in middleware scenarios)
		if c == nil {
			log.Println("fiber.Ctx is nil")
			return fiber.ErrInternalServerError
		}
		animals, err := animal_service.GetAllAnimals()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to get all animals")
		}
		return c.JSON(animals)
	}
}
