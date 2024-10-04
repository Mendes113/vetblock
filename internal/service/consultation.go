package service

import (
	"context"
	"errors"
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
	consultations, err := repo.FindConsultationByVeterinaryCRVM(context.Background(), crvm)
	if err != nil {
		return nil, err
	}

	// Verifica se não há consultas
	if len(consultations) == 0 {
		return nil, errors.New("nenhuma consulta encontrada")
	}

	// Encontra a próxima consulta (a mais próxima da data atual)
	var nextConsultation *model.Consultation
	now := time.Now()

	for _, consultation := range consultations {
		if consultation.ConsultationDate.After(now) {
			if nextConsultation == nil || consultation.ConsultationDate.Before(nextConsultation.ConsultationDate) {
				nextConsultation = &consultation
			}
		}
	}

	// Se não houver consulta futura, retorna erro
	if nextConsultation == nil {
		return nil, errors.New("nenhuma consulta futura encontrada")
	}

	return nextConsultation, nil
}
