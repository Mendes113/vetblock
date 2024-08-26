package handlers

import (
	"strconv"
	"time"
	"vetblock/internal/db/model"
	"vetblock/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MedicationResponse struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name" validate:"required,min=2,max=100"`
	ActivePrinciples []string  `json:"active_principle" validate:"required"`
	Manufacturer     string    `json:"manufacturer" validate:"required"`
	Concentration    string    `json:"concentration" validate:"required"`
	Presentation     string    `json:"presentation" validate:"required"`
	Quantity         int       `json:"quantity" validate:"required,gte=0"`
	ExpirationDate   string    `json:"expiration_date" validate:"required"`
}

func AddMedicationHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var medicationResponse MedicationResponse

		// Parse JSON body
		if err := c.BodyParser(&medicationResponse); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body: " + err.Error())
		}

		// Convert string to time.Time for ExpirationDate
		expiration, err := time.Parse("2006-01-02", medicationResponse.ExpirationDate)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid date format for expiration_date: " + err.Error())
		}

		// Convert MedicationResponse to Medication model
		medicationModel := model.Medication{
			ID:               medicationResponse.ID,
			Name:             medicationResponse.Name,
			ActivePrinciples: medicationResponse.ActivePrinciples, // Atualize o nome do campo para ActivePrinciples
			Manufacturer:     medicationResponse.Manufacturer,
			Concentration:    medicationResponse.Concentration,
			Presentation:     medicationResponse.Presentation,
			Quantity:         medicationResponse.Quantity,
			Expiration:       expiration,
			// Campos adicionais, se necessário
		}

		// Add medication to the database
		err = service.AddMedication(medicationModel)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to add medication: " + err.Error())
		}

		// Return the added medication response
		return c.Status(fiber.StatusCreated).JSON(medicationResponse)
	}
}
func GetMedicationByIDHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
		}

		medication, err := service.GetMedicationByID(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Medication not found")
		}

		return c.Status(fiber.StatusOK).JSON(medication)
	}
}

func DeleteMedicationHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		msg, err := service.DeleteMedication(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete medication")
		}

		return c.Status(fiber.StatusOK).SendString(msg)
	}
}

func UpdateMedicationHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var medication model.Medication
		if err := c.BodyParser(&medication); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
		}

		if err := service.UpdateMedication(medication); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to update medication")
		}

		return c.Status(fiber.StatusOK).JSON(medication)
	}
}

func GetAllMedicationsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		medications, err := service.GetAllMedications()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve medications")
		}

		return c.Status(fiber.StatusOK).JSON(medications)
	}
}

func GetMedicationClosestExpirationDateHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		medication, err := service.GetMedicationClosestExpirationDate()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve medication with closest expiration date")
		}

		return c.Status(fiber.StatusOK).JSON(medication)
	}
}

func GetExpiredMedicationsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		medications, err := service.GetExpiredMedications()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve expired medications")
		}

		return c.Status(fiber.StatusOK).JSON(medications)
	}
}

func GetMedicationsWillExpireInDaysHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		days, err := strconv.Atoi(c.Params("days"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid days format")
		}

		medications, err := service.GetMedicationsWillExpireInDays(days)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve medications")
		}

		return c.Status(fiber.StatusOK).JSON(medications)
	}
}

func GetMedicationByBatchNumberHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		batchNumber := c.Params("batch_number")

		medication, err := service.GetMedicationByBatchNumber(batchNumber)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Medication not found with given batch number")
		}

		return c.Status(fiber.StatusOK).JSON(medication)
	}
}

func GetMedicationByNameHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")

		medication, err := service.GetMedicationByName(name)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Medication not found with given name")
		}

		return c.Status(fiber.StatusOK).JSON(medication)
	}
}

func GetMedicationByActiveSubstanceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		activeSubstance := c.Params("active_substance")

		medications, err := service.GetMedicationByActiveSubstance(activeSubstance)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Medication not found with the given active substance")
		}

		return c.Status(fiber.StatusOK).JSON(medications)
	}
}
