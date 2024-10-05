package service

import (
	"context"
	"errors"
	"log"
	"time"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"

	"github.com/google/uuid"
)

// Função para verificar se a consulta já existe e retornar erro se necessário
func checkConsultationExistence(repo repository.ConsultationRepository, id uuid.UUID) (*model.Consultation, error) {
	existingConsultation, err := repo.FindConsultationByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if existingConsultation == nil {
		return nil, errors.New("consulta não encontrada")
	}
	return existingConsultation, nil
}


func findConflictingConsultations(repo repository.ConsultationRepository, consultation *model.Consultation) ([]model.Consultation, error) {
	log.Print("Finding conflicting consultations")
	
	// Formato de data e hora utilizado
	const dateTimeLayout = "2006-01-02 15:04"

	// Combina a data e a hora da nova consulta (conversão de string para time.Time)
	consultationDateTimeStr := consultation.ConsultationDate.Format("2006-01-02") + " " + consultation.ConsultationHour
	consultationDateTime, err := time.Parse(dateTimeLayout, consultationDateTimeStr)
	if err != nil {
		return nil, err
	}

	// Busque consultas pela data
	consultationDateStr := consultation.ConsultationDate.Format("2006-01-02")
	consultations, err := repo.FindConsultationByDate(context.Background(), consultationDateStr)
	if err != nil {
		return nil, err
	}

	var conflictingConsultations []model.Consultation
	for _, c := range consultations {
		// Verifique se a consulta é diferente da que está sendo adicionada
		if c.ID != consultation.ID {
			// Combina a data e a hora da consulta existente (conversão de string para time.Time)
			existingConsultationDateTimeStr := c.ConsultationDate.Format("2006-01-02") + " " + c.ConsultationHour
			existingConsultationDateTime, err := time.Parse(dateTimeLayout, existingConsultationDateTimeStr)
			if err != nil {
				return nil, err
			}

			// Calcula a diferença de tempo entre as consultas em minutos
			diff := consultationDateTime.Sub(existingConsultationDateTime).Minutes()

			// Verifica se a diferença é menor que 15 minutos (ou seja, há um conflito)
			if diff < 15 && diff > -15 {
				conflictingConsultations = append(conflictingConsultations, c)
			}
		}
	}
	log.Print("Found conflicting consultations:", conflictingConsultations)
	return conflictingConsultations, nil
}

func AddConsultation(repo repository.ConsultationRepository, consultation *model.Consultation, getVetFunc func(string) (*model.Veterinary, error), getAnimalFunc func(uuid.UUID) (*model.Animal, error)) error {
	// Verifique se a consulta já existe
	existingConsultation, _ := repo.FindConsultationByID(context.Background(), consultation.ID)
	if existingConsultation != nil {
		return errors.New("consulta já existe")
	}

	// Verifique se o veterinário existe
	vet, err := getVetFunc(consultation.CRVM)
	if err != nil {
		return err
	}
	if vet == nil {
		return errors.New("veterinário não encontrado")
	}

	// Verifique se o animal existe
	animal, err := getAnimalFunc(consultation.AnimalID)
	if err != nil {
		return err
	}
	if animal == nil {
		return errors.New("animal não encontrado")
	}

	// Verifique conflitos de horário
	conflictingConsultations, err := findConflictingConsultations(repo, consultation)
	if err != nil {
		return err
	}
	if len(conflictingConsultations) > 0 {
		return errors.New("já existe uma consulta no mesmo horário")
	}

	// Salve a nova consulta
	if err := repo.SaveConsultation(context.Background(), consultation); err != nil {
		return err
	}

	return nil
}


func UpdateConsultation(repo repository.ConsultationRepository, id uuid.UUID, updatedConsultation *model.Consultation) error {
	consultation, err := checkConsultationExistence(repo, id)
	if err != nil {
		return err
	}

	// Atualiza os campos da consulta
	consultation.AnimalID = updatedConsultation.AnimalID
	consultation.CRVM = updatedConsultation.CRVM
	consultation.ConsultationDate = updatedConsultation.ConsultationDate
	consultation.ConsultationDescription = updatedConsultation.ConsultationDescription
	consultation.ConsultationType = updatedConsultation.ConsultationType
	consultation.ConsultationPrescription = updatedConsultation.ConsultationPrescription
	consultation.ConsultationPrice = updatedConsultation.ConsultationPrice

	if err := repo.SaveConsultation(context.Background(), consultation); err != nil {
		return err
	}

	return nil
}

func DeleteConsultation(repo repository.ConsultationRepository, id uuid.UUID) error {
	consultation, err := checkConsultationExistence(repo, id)
	if err != nil {
		return err
	}

	if err := repo.DeleteConsultation(context.Background(), consultation.ID); err != nil {
		return err
	}

	return nil
}
func GetNextConsultationByVeterinaryCRVM(repo repository.ConsultationRepository, crvm string) (*model.Consultation, error) {
	// Busca todas as consultas relacionadas ao CRVM do veterinário
	consultations, err := repo.FindConsultationByVeterinaryCRVM(context.Background(), crvm)
	if err != nil {
		log.Printf("Erro ao buscar consultas para o CRVM %s: %v", crvm, err)
		return nil, err
	}
	log.Printf("Consultas encontradas para o CRVM %s: %v", crvm, consultations)

	// Verifica se não há consultas
	if len(consultations) == 0 {
		log.Printf("Nenhuma consulta encontrada para o CRVM %s", crvm)
		return nil, errors.New("nenhuma consulta encontrada")
	}

	// Data e hora atual em UTC
	now := time.Now().UTC()
	log.Printf("Data e hora atuais (UTC): %v", now)

	var nextConsultation *model.Consultation

	// Itera pelas consultas encontradas
	for _, consultation := range consultations {
		// Construa a data e hora da consulta completa
		consultationDateTime := consultation.ConsultationDate.UTC()
		if consultation.ConsultationHour != "" {
			consultationDateTime, err = time.Parse("2006-01-02 15:04", consultation.ConsultationDate.Format("2006-01-02")+" "+consultation.ConsultationHour)
			if err != nil {
				log.Printf("Erro ao interpretar data e hora da consulta: %v", err)
				return nil, errors.New("erro ao interpretar data e hora da consulta")
			}
			consultationDateTime = consultationDateTime.UTC()
		}
		log.Printf("Data e hora da consulta (UTC): %v", consultationDateTime)

		// Verifica se a consulta é após a data e hora atuais
		if consultationDateTime.After(now) {
			log.Printf("Consulta futura encontrada: %v", consultationDateTime)

			// Se ainda não temos uma próxima consulta ou a consulta atual for mais próxima
			if nextConsultation == nil || consultationDateTime.Before(nextConsultation.ConsultationDate.UTC()) {
				log.Printf("Atualizando próxima consulta para: %v", consultationDateTime)
				// Atualiza a próxima consulta
				nextConsultation = &consultation
				nextConsultation.ConsultationDate = model.CustomDate{consultationDateTime}
			}
		} else {
			log.Printf("Consulta anterior ou no mesmo horário que o atual: %v", consultationDateTime)
		}
	}

	// Se não houver consulta futura, retorna erro
	if nextConsultation == nil {
		log.Printf("Nenhuma consulta futura encontrada para o CRVM %s", crvm)
		return nil, errors.New("nenhuma consulta futura encontrada")
	}

	log.Printf("Próxima consulta encontrada: %v", nextConsultation)
	return nextConsultation, nil
}
