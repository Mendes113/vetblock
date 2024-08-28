package service

import (
	"errors"
	"log"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"

	"github.com/google/uuid"
)

func getMedicationRepo() *repository.MedicationRepository {
	return repository.NewMedicationRepository()
}

func AddMedication(medication *model.Medication) (*model.Medication,error) {
    log.Println("adding medication transaction")
    log.Print("existingMedication")
    log.Print(medication)

    repo := repository.NewMedicationRepository()
    existingMedication := repo.FindByUniqueAttributes(medication)
   

    if existingMedication != nil {
        if err := repo.IncreaseMedicationQuantity(existingMedication.ID, medication.Quantity); err != nil {
            log.Print("Error increasing medication quantity:", err)
            return nil,err
        }
        log.Print("Medication already exists, increasing quantity")
        return existingMedication, nil
    }

    if err := repo.SaveMedication(medication); err != nil {
        log.Print("Error saving new medication:", err)
        return nil, err
    }

    log.Print("New medication added successfully")
    return medication, nil
}


func GetMedicationByID(id uuid.UUID) (*model.Medication, error) {
	repo := getMedicationRepo()
	medication, err := repo.FindMedicationByID(id)
	if err != nil {
		return nil, err
	}

	return medication, nil
}

func DeleteMedication(id string) (string, error) {
	repo := getMedicationRepo()
	msg, err := repo.DeleteMedication(id)
	if err != nil {
		return "", err
	}

	return msg, nil
}

func UpdateMedication(medication model.Medication) error {
	repo := getMedicationRepo()
	medicationFound, err := repo.FindMedicationByID(medication.ID)
	if err != nil {
		return err
	}
	if medicationFound == nil {
		return errors.New("medication não encontrada")
	}

	if err := repo.SaveMedication(&medication); err != nil {
		return err
	}

	return nil
}

func GetAllMedications() ([]model.Medication, error) {
	repo := getMedicationRepo()
	medications, err := repo.FindAllMedications()
	if err != nil {
		return nil, err
	}

	return medications, nil
}

func GetMedicationClosestExpirationDate() (*model.Medication, error) {
	repo := getMedicationRepo()
	medication, err := repo.FindMedicationClosestExpirationDate()
	if err != nil {
		return nil, err
	}

	return medication, nil
}

func CheckMedicationExistence(repo repository.MedicationRepository, id uuid.UUID) (*model.Medication, error) {
	medication, err := repo.FindMedicationByID(id)
	if err != nil {
		return nil, err
	}
	if medication == nil {
		return nil, errors.New("medication não encontrada")
	}
	return medication, nil
}


//already expired medications
func GetExpiredMedications() ([]model.Medication, error) {
	repo := getMedicationRepo()
	medication, err := repo.FindMedicationExpired() 
	if err != nil {
		return nil, err
	}

	
	medications := []model.Medication{}
	if medication != nil { 
		medications = append(medications, *medication)
	}

	return medications, nil 
}


//medications that will expire in the next 30 days
func GetMedicationsWillExpireInDays(days int) ([]model.Medication, error) {
	repo := getMedicationRepo()
	medications, err := repo.FindMedicationWillExpireInDays(days) 
	if err != nil {
		return nil, err
	}

	return medications, nil 
}

//find medications by batch number
func GetMedicationByBatchNumber(batchNumber string) (*model.Medication, error) {
	repo := getMedicationRepo()
	medication, err := repo.FindMedicationByBatchNumber(batchNumber)
	if err != nil {
		return nil, err
	}

	return medication, nil
}

//find medications by name
func GetMedicationByName(name string) (*model.Medication, error) {
	repo := getMedicationRepo()
	medication, err := repo.FindMedicationByName(name)
	if err != nil {
		return nil, err
	}

	return medication, nil
}

//find medications by name
func GetMedicationByActiveSubstance(activeSubstance string) (*model.Medication, error) {
	repo := getMedicationRepo()
	medications, err := repo.FindMedicationByActiveSubstance(activeSubstance)
	if err != nil {
		return nil, err
	}

	if len(medications) > 0 {
		return &medications[0], nil // Retorna o primeiro medicamento encontrado
	}
	errorMessage := "medicamento não encontrado com a substância ativa informada" + activeSubstance
	return nil, errors.New(errorMessage) 
}


