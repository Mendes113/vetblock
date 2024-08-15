package handlers

import (
	"log"
	"strconv"
	"vetblock/internal/db/model"
	"vetblock/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Handler para agendar uma consulta
func ScheduleConsultationHandler(c *fiber.Ctx) error {
	var consultation model.Consultation

	if err := c.BodyParser(&consultation); err != nil {
		log.Printf("Erro ao analisar a solicitação: %v", err)
		log.Print(&consultation)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Solicitação inválida",
		})
	}

	consultation.ID = uuid.New()
	sender := c.Query("sender")
	receiver := c.Query("receiver")
	amount := c.QueryFloat("amount", 0)

	if err := service.ScheduleConsultation(consultation, sender, receiver, amount); err != nil {
		log.Printf("Erro ao agendar consulta: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Não foi possível agendar a consulta",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(consultation)
}

// Handler para cancelar uma consulta
func CancelConsultationHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	consultationID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	consultation, err := service.GetConsultationByID(consultationID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar consulta",
		})
	}
	if consultation == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Consulta não encontrada",
		})
	}

	sender := c.Query("sender")
	receiver := c.Query("receiver")
	amount := c.QueryFloat("amount", 0)

	if err := service.CancelConsultation(*consultation, sender, receiver, amount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Não foi possível cancelar a consulta",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Consulta cancelada com sucesso",
	})
}

// Handler para confirmar uma consulta
func ConfirmConsultationHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	consultationID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	consultation, err := service.GetConsultationByID(consultationID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar consulta",
		})
	}
	if consultation == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Consulta não encontrada",
		})
	}

	sender := c.Query("sender")
	receiver := c.Query("receiver")
	amount := c.QueryFloat("amount", 0)

	if err := service.ConfirmConsultation(*consultation, sender, receiver, amount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Não foi possível confirmar a consulta",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Consulta confirmada com sucesso",
	})
}

// Handler para atualizar uma consulta
func UpdateConsultationHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	consultationID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	consultation, err := service.GetConsultationByID(consultationID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar consulta",
		})
	}
	if consultation == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Consulta não encontrada",
		})
	}

	if err := c.BodyParser(&consultation); err != nil {
		log.Printf("Erro ao analisar a solicitação: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Solicitação inválida",
		})
	}

	sender := c.Query("sender")
	receiver := c.Query("receiver")
	amount := c.QueryFloat("amount", 0)

	if err := service.UpdateConsultation(*consultation, sender, receiver, amount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Não foi possível atualizar a consulta",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Consulta atualizada com sucesso",
	})
}

// Handler para buscar consultas por ID do animal
func GetConsultationByAnimalIDHandler(c *fiber.Ctx) error {
	animalID := c.Params("animal_id")

	
	consultationID, err := uuid.Parse(animalID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID de Animal inválido",
		})
	}

	consultations, err := service.GetConsultationByAnimalID(consultationID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar consultas",
		})
	}

	return c.JSON(consultations)
}

// Handler para buscar consultas por ID do veterinário
func GetConsultationByVeterinaryIDHandler(c *fiber.Ctx) error {
	crvm := c.Params("crvm")
	crvmInt, err := strconv.Atoi(crvm)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "CRVM inválido",
		})
	}
	
	consultations, err := service.GetConsultationByVeterinaryCRVM(crvmInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar consultas",
		})
	}

	return c.JSON(consultations)
}
