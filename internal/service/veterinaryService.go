package service

import (
	"log"
	"time"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"
)

func getVeterinaryRepo() *repository.VeterinaryRepository {
	return repository.NewVeterinaryRepository()
}

func GetVeterinaryByCRVM(crvm string) (*model.Veterinary, error) {
	repo := getVeterinaryRepo()
	return repo.FindVeterinaryByCRVM(crvm)
}

func AddVeterinary(veterinary model.Veterinary) error {

	veterinary.CreatedAt = time.Now()
	veterinary.UpdatedAt = time.Now()
	

	log.Println("adding veterinary transaction")
	repo := getVeterinaryRepo()
	return repo.SaveVeterinary(veterinary)
}
