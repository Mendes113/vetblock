package handlers

import (
    "time"
    "vetblock/internal/db/model"
    "vetblock/internal/service"
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "github.com/go-playground/validator/v10"
)

// Defina o validador global
var validate = validator.New()

type ConsultationRequest struct {
    ID               uuid.UUID `json:"id"`
    AnimalID         uuid.UUID `json:"animal_id" validate:"required"`
    VeterinaryCRVM   string    `json:"crvm" validate:"required"`
    ConsultationDate string    `json:"consultation_date" validate:"required"`
    Reason           string    `json:"reason" validate:"required"`
    Consultation_Type string    `json:"consultation_type"`
    Consultation_Hour string    `json:"consultation_hour"`
    Consultation_Prescription string    `json:"consultation_prescription"`
    Consultation_Status string    `json:"consultation_status"`
    Observation      string    `json:"observation"`
    Consultation_Price float64    `json:"consultation_price"`
}

func AddConsultationHandler() fiber.Handler {
    return func(c *fiber.Ctx) error {
        var consultation ConsultationRequest
        consultation.ID = uuid.New()

        // Parse o corpo da solicitação
        if err := c.BodyParser(&consultation); err != nil {
            return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
        }

        // Valide os dados da consulta
        if err := validate.Struct(&consultation); err != nil {
            errs := err.(validator.ValidationErrors)
            var errorMessages []string
            for _, e := range errs {
                errorMessages = append(errorMessages, e.Field()+" is "+e.Tag())
            }
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": errorMessages,
            })
        }

        // Valide o formato da data
        parsedDate, err := time.Parse("2006-01-02", consultation.ConsultationDate)
        if err != nil {
            return c.Status(fiber.StatusBadRequest).SendString("Invalid date format")
        }

        // Converta o ConsultationResponse para Consultation
        consultationModel := model.Consultation{
            ID:               consultation.ID,
            AnimalID:         consultation.AnimalID,
            CRVM:             consultation.VeterinaryCRVM,
            ConsultationDate: parsedDate,
            Reason:           consultation.Reason,
            Observation:      consultation.Observation,
            ConsultationType: consultation.Consultation_Type,
            ConsultationHour: consultation.Consultation_Hour,
            ConsultationPrescription: consultation.Consultation_Prescription,
            ConsultationStatus: consultation.Consultation_Status,
            ConsultationPrice: consultation.Consultation_Price,

        }

        err = service.AddConsultation(&consultationModel)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString("Failed to add consultation transaction")
        }

        return c.Status(fiber.StatusCreated).JSON(fiber.Map{
            "message": "Consulta adicionada com sucesso",
        })
        
    }
}

//get next vet(using crvm) consultation
func GetNextConsultationHandler() fiber.Handler {
    return func(c *fiber.Ctx) error {
        crvm := c.Params("crvm")

        consultation, err := service.GetNextConsultationByVeterinaryCRVM(crvm)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": err.Error(),
            })
        }

        return c.Status(fiber.StatusOK).JSON(consultation)
    }
}