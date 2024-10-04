package service

import (
	"context"
	"errors"
	"log"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"

	"github.com/google/uuid"
)

// type ConsultationService struct{
// 	repo repository.ConsultationRepository
// }

// Função para obter o repositório de consultas
func getConsultationRepo() *repository.ConsultationRepository {
	return repository.NewConsultationRepository(
		repository.GetDB(),
	)
}

// Função para verificar se a consulta já existe e retornar erro se necessário
func checkConsultationExistence(repo *repository.ConsultationRepository, id uuid.UUID) (*model.Consultation, error) {
	existingConsultation, err := repo.FindConsultationByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if existingConsultation == nil {
		return nil, errors.New("consulta não encontrada")
	}
	return existingConsultation, nil
}

func AddConsultation(consultation *model.Consultation) error {
	repo := getConsultationRepo()
	log.Print(consultation) // Você pode querer verificar o conteúdo aqui

	// Verifique se a consulta já existe
	existingConsultation, _ := repo.FindConsultationByID(context.Background(), consultation.ID)
	
	if existingConsultation != nil {
		return errors.New("consulta já existe")
	}

	// Salve a nova consulta
	if err := repo.SaveConsultation(context.Background(), consultation); err != nil {
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

func UpdateConsultation(id uuid.UUID, updatedConsultation *model.Consultation) error {
	repo := getConsultationRepo()
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

func DeleteConsultation(id uuid.UUID) error {
	repo := getConsultationRepo()
	consultation, err := checkConsultationExistence(repo, id)
	if err != nil {
		return err
	}

	if err := repo.DeleteConsultation(context.Background(), consultation.ID); err != nil {
		return err
	}

	return nil
}

// Adapte a função getConsultationBy para aceitar métodos com receiver de ponteiro
func getConsultationBy(
	findFunc func(ctx context.Context, id uuid.UUID) (*model.Consultation, error),
	ctx context.Context,
	id uuid.UUID,
) (*model.Consultation, error) {
	return findFunc(ctx, id)
}

func GetConsultationByID(id uuid.UUID) (*model.Consultation, error) {
	repo := getConsultationRepo() // Crie uma instância do repositório
	return getConsultationBy(repo.FindConsultationByID, context.Background(), id)
}

func GetConsultationByAnimalID(animalID uuid.UUID) ([]model.Consultation, error) {
	repo := getConsultationRepo()
	return repo.FindConsultationByAnimalID(context.Background(), animalID)
}

func GetConsultationByVeterinaryCRVM(crvm string) ([]model.Consultation, error) {
	repo := getConsultationRepo()
	return repo.FindConsultationByVeterinaryCRVM(context.Background(), crvm)
}

func GetConsultationByDate(date string) ([]model.Consultation, error) {
	repo := getConsultationRepo()
	return repo.FindConsultationByDate(context.Background(), date)
}

func GetConsultationByDateRange(startDate, endDate string) ([]model.Consultation, error) {
	repo := getConsultationRepo()
	return repo.FindConsultationByDateRange(context.Background(), startDate, endDate)
}

func GetConsultationByAnimalIDAndDateRange(animalID uuid.UUID, startDate, endDate string) ([]model.Consultation, error) {
	repo := getConsultationRepo()
	return repo.FindConsultationByAnimalIDAndDateRange(context.Background(), animalID, startDate, endDate)
}

//next vet consultation
func GetNextConsultationByVeterinaryCRVM(crvm string) (*model.Consultation, error) {
	repo := getConsultationRepo()
	
	consultations, err := repo.FindConsultationByVeterinaryCRVM(context.Background(), crvm)
	if err != nil {
		return nil, err
	}

	if len(consultations) == 0 {
		return nil, nil
	}

	// Encontre a próxima consulta
	nextConsultation := consultations[0]
	for _, consultation := range consultations {
		if consultation.ConsultationDate.After(nextConsultation.ConsultationDate) {
			nextConsultation = consultation
		}
	}

	return &nextConsultation, nil
}

func GetConsultationByAnimalIDAndDate(animalID uuid.UUID, date string) (*model.Consultation, error) {
	repo := getConsultationRepo()
	consultations, err := repo.FindConsultationByAnimalIDAndDate(context.Background(), animalID, date)
	if err != nil {
		return nil, err
	}
	if len(consultations) == 0 {
		return nil, nil
	}
	return &consultations[0], nil
}