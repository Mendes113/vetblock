package handlers

import (
	"context"
	"log"
	"strings"
	"time"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"
	"vetblock/internal/service"

	"github.com/go-playground/validator/v10"
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

type DosageResponse struct {
	ID                uuid.UUID        `json:"id"`
	AnimalID          uuid.UUID        `json:"animal_id" validate:"required,uuid"`
	MedicationID      uuid.UUID        `json:"medication_id" validate:"required,uuid"`
	StartDate         model.CustomDate `json:"start_date" validate:"required"`                 // Usa CustomDate
	EndDate           model.CustomDate `json:"end_date" validate:"required,gtfield=StartDate"` // Usa CustomDate
	Quantity          int              `json:"quantity" validate:"gte=0"`
	Dosage            string           `json:"dosage" validate:"required"`
	ConsultationID    *uuid.UUID       `json:"consultation_id"`    // Relacionamento opcional
	HospitalizationID *uuid.UUID       `json:"hospitalization_id"` // Relacionamento opcional
}

func AddAnimalHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var animal AnimalResponse

		if err := c.BodyParser(&animal); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		//validate cpf(clean)
		animal.CPFTutor = CleanCpf(animal.CPFTutor)

		// Convert the AnimalResponse to Animal
		animalModel := model.Animal{
			ID:          uuid.New(),
			Name:        animal.Name,
			Species:     animal.Species,
			Breed:       animal.Breed,
			Weight:      animal.Weight,
			Age:         animal.Age,
			Description: animal.Description,
			CPFTutor:    animal.CPFTutor,
		}

		// Validação dos campos do Animal
		if err := service.ValidateAnimal(animalModel); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		// Adiciona o animal usando o serviço
		err := animal_service.AddAnimal(animalModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(animal)
	}
}

func CleanCpf(cpf string) string {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	return cpf
}

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

		// Atualiza o animal
		if err := animal_service.UpdateAnimal(id, animal); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to update animal")
		}

		return c.Status(fiber.StatusOK).JSON(animal)
	}
}

func DeleteAnimalHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
		}

		// Deleta o animal
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

// Handler para adicionar uma nova dosagem
func AddDosageHandler(dosageService *service.DosageService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var dosage DosageResponse

		// Parse o corpo da requisição
		if err := c.BodyParser(&dosage); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		// Valida o corpo da requisição usando validator
		if err := validate.Struct(&dosage); err != nil {
			errs := err.(validator.ValidationErrors)
			var errorMessages []string
			for _, e := range errs {
				errorMessages = append(errorMessages, e.Field()+" is "+e.Tag())
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": errorMessages,
			})
		}

		// Verifica se o animal existe
		animal, err := animal_service.GetAnimalByID(dosage.AnimalID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving animal")
		}
		if animal == nil {
			return c.Status(fiber.StatusNotFound).SendString("Animal not found")
		}

		// Gera um novo UUID para a dosagem
		dosageID := uuid.New()

		// Converte DosageResponse para o model Dosage
		dosageModel := model.Dosage{
			ID:                dosageID,
			AnimalID:          dosage.AnimalID,
			MedicationID:      dosage.MedicationID,
			StartDate:         model.CustomDate{Time: time.Now().Truncate(24 * time.Hour)},          // Truncando hora para manter só a data
			EndDate:           model.CustomDate{Time: dosage.EndDate.Time.Truncate(24 * time.Hour)}, // Truncando hora para manter só a data
			Quantity:          dosage.Quantity,
			Dosage:            dosage.Dosage,
			ConsultationID:    nilIfEmpty(dosage.ConsultationID),
			HospitalizationID: nilIfEmpty(dosage.HospitalizationID),
		}

		// Chama o serviço para adicionar a dosagem
		err = dosageService.AddDosage(context.Background(), &dosageModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to add dosage transaction")
		}

		// Retorna a dosagem criada
		dosage.ID = dosageID
		return c.Status(fiber.StatusCreated).JSON(dosage)
	}
}

// nilIfEmpty é uma função auxiliar para lidar com campos opcionais
func nilIfEmpty(id *uuid.UUID) *uuid.UUID {
	if id == nil || *id == uuid.Nil {
		return nil
	}
	return id
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

func GetAnimalByIDHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
		}

		animal, err := animal_service.GetAnimalByID(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to get animal")
		}
		if animal == nil {
			return c.Status(fiber.StatusNotFound).SendString("Animal not found")
		}

		return c.JSON(animal)
	}
}
