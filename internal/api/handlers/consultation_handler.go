package handlers

import (
	"fmt"
	"log"
	"time"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"
	"vetblock/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func AddConsultationHandler(repo repository.ConsultationRepository) fiber.Handler {
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

        parsedDate, err := time.Parse("2006-01-02", consultation.ConsultationDate)

        // Valide o formato da hora (se fornecido)
        if consultation.Consultation_Hour != "" {
            _, err := time.Parse("15:04", consultation.Consultation_Hour)
            if err != nil {
                return c.Status(fiber.StatusBadRequest).SendString("Invalid time format")
            }
        }

        // Converta o ConsultationRequest para o modelo Consultation
        consultationModel := model.Consultation{
            ID:               consultation.ID,
            AnimalID:         consultation.AnimalID,
            CRVM:             consultation.VeterinaryCRVM,
            ConsultationDate: model.CustomDate{Time: parsedDate},
            Reason:           consultation.Reason,
            Observation:      consultation.Observation,
            ConsultationType: consultation.Consultation_Type,
            ConsultationHour: consultation.Consultation_Hour,
            ConsultationPrescription: consultation.Consultation_Prescription,
            ConsultationStatus:     consultation.Consultation_Status,
            ConsultationPrice:      consultation.Consultation_Price,
        }

        // Função para buscar o veterinário pelo CRVM
        getVeterinaryByCRVM := func(crvm string) (*model.Veterinary, error) {
            vet, err := service.GetVeterinaryByCRVM(crvm)
            if err != nil {
                return nil, err
            }
            return vet, nil
        }

        // Função para buscar o animal pelo ID
        getAnimalByID := func(animalID uuid.UUID) (*model.Animal, error) {
            animal, err := animal_service.GetAnimalByID(animalID)
            if err != nil {
                return nil, err
            }
            return animal, nil
        }

        // Chame a função AddConsultation com todos os parâmetros necessários
        err = service.AddConsultation(repo, &consultationModel, getVeterinaryByCRVM, getAnimalByID)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString("Failed to add consultation transaction")
        }

        return c.Status(fiber.StatusCreated).JSON(fiber.Map{
            "message": fmt.Sprintf("Consulta adicionada com sucesso para o dia %s às %s.", consultation.ConsultationDate, consultation.Consultation_Hour),
        })
    }
}



//get next vet(using crvm) consultation
func GetNextConsultationHandler(repo repository.ConsultationRepository) fiber.Handler {
    return func(c *fiber.Ctx) error {
        crvm := c.Params("crvm")
        log.Println("CRVM: ", crvm)
        consultation, err := service.GetNextConsultationByVeterinaryCRVM(repo, crvm)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "message": err.Error(),
            })
        }

        return c.Status(fiber.StatusOK).JSON(consultation)
    }
}