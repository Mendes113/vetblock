package service

import (
	"errors"
	"log"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"
)

// Função para obter o repositório de consultas
func getConsultationRepo() *repository.ConsultationRepository {
	return repository.NewConsultationRepository()
}

// Função para verificar se a consulta já existe e retornar erro se necessário
func checkConsultationExistence(repo repository.ConsultationRepository, consultation model.Consultation) (*model.Consultation, error) {
	existingConsultation, err := repo.FindConsultationByID(consultation)
	if err != nil {
		return nil, err
	}
	if existingConsultation == nil {
		return nil, errors.New("consulta não encontrada")
	}
	return existingConsultation, nil
}

func AddConsultation(consultation model.Consultation) error {
	repo := getConsultationRepo()
	log.Print(consultation) // Você pode querer verificar o conteúdo aqui

	// Certifique-se de que FindConsultationByID usa ponteiros
	existingConsultation, err := repo.FindConsultationByID(consultation)
	if err != nil {
		return err
	}
	log.Print("existingConsultation")
	log.Print(existingConsultation)
	if existingConsultation != nil {
		return errors.New("consulta já existe")
	}

	if err := repo.SaveConsultation(&consultation); err != nil {
		return err
	}

	// Verifique se o veterinário existe
	vet, err := GetVeterinaryByCRVM(consultation.CRVM)
	if err != nil {
		return err
	}

	if vet == nil {
		return errors.New("veterinário não encontrado")
	}

	// Verifique se o animal existe
	animal, err := GetAnimalByID(consultation.AnimalID)
	if err != nil {
		return err
	}

	if animal == nil {
		return errors.New("animal não encontrado")
	}

	return nil
}

func UpdateConsultation(id model.Consultation, updatedConsultation model.Consultation) error {
	repo := getConsultationRepo()
	consultation, err := checkConsultationExistence(*repo, id)
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

	if err := repo.SaveConsultation(consultation); err != nil {
		return err
	}

	return nil
}

func DeleteConsultation(id model.Consultation) error {
	repo := getConsultationRepo()
	consultation, err := checkConsultationExistence(*repo, id)
	if err != nil {
		return err
	}

	if _, err := repo.DeleteConsultation(*consultation); err != nil {
		return err
	}

	return nil
}

// Adapte a função getConsultationBy para aceitar métodos com receiver de ponteiro
func getConsultationBy(criteria func(*repository.ConsultationRepository, model.Consultation) (*model.Consultation, error), id model.Consultation) (*model.Consultation, error) {
	repo := getConsultationRepo()
	return criteria(repo, id)
}

func GetConsultationByID(id model.Consultation) (*model.Consultation, error) {
	return getConsultationBy((*repository.ConsultationRepository).FindConsultationByID, id)
}

func GetConsultationByAnimalID(id model.Consultation) (*model.Consultation, error) {
	return getConsultationBy((*repository.ConsultationRepository).FindConsultationByAnimalID, id)
}

func GetConsultationByVeterinaryCRVM(id model.Consultation) (*model.Consultation, error) {
	return getConsultationBy((*repository.ConsultationRepository).FindConsultationByVeterinaryCRVM, id)
}

func GetConsultationByDate(id model.Consultation) (*model.Consultation, error) {
	return getConsultationBy((*repository.ConsultationRepository).FindConsultationByDate, id)
}

func GetConsultationByDateRange(id model.Consultation) (*model.Consultation, error) {
	return getConsultationBy((*repository.ConsultationRepository).FindConsultationByDateRange, id)
}

func GetConsultationByAnimalIDAndDateRange(id model.Consultation) (*model.Consultation, error) {
	return getConsultationBy((*repository.ConsultationRepository).FindConsultationByAnimalIDAndDateRange, id)
}
