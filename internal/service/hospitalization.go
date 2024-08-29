package service

import (
	// "encoding/json"
	"errors"
	// "log"
	"time"
	"vetblock/internal/db/model"
	"vetblock/internal/db/repository"

	"github.com/google/uuid"
)

const (
	ErrInvalidHospitalizationID = "id da hospitalização inválido"
	ErrInvalidPatientID         = "id do paciente inválido"
	ErrInvalidStartDate         = "data de início inválida"
	ErrInvalidEndDate           = "data de término inválida"
	ErrInvalidReason            = "motivo inválido"
	ErrInvalidDoctorID          = "id do médico inválido"
	ErrInvalidMedications       = "lista de medicamentos inválida"
)

func gethospitalizationRepo() *repository.HospitalizationRepository {
	return repository.NewHospitalizationRepository()
}

func ValidateHospitalization(hospitalization model.Hospitalization) error {

	//validate all medications
	for _, medication := range hospitalization.Medications {
		if err := verifyMedication(medication); err != nil {
			return err
		}
	}
	

	validations := []struct {
		condition bool
		errMsg    string
	}{
		{hospitalization.ID == uuid.UUID{}, ErrInvalidHospitalizationID},
		{hospitalization.PatientID == uuid.UUID{}, ErrInvalidPatientID},
		{hospitalization.StartDate == time.Time{}, ErrInvalidStartDate},
		{hospitalization.EndDate == time.Time{}, ErrInvalidEndDate},
		{hospitalization.Reason == "", ErrInvalidReason},
		{hospitalization.CRVM == 0, ErrInvalidDoctorID},
		{len(hospitalization.Medications) == 0, ErrInvalidMedications},
	}

	for _, v := range validations {
		if v.condition {
			return errors.New(v.errMsg)
		}
	}

	return nil
}

//verify if the medication exists
func verifyMedication(medicationName string) error {
	repo := repository.NewMedicationRepository()
	_, err := repo.FindMedicationByName(medicationName)
	if err != nil {
		return err
	}
	return errors.New("medication not found")
}
	


func AddHospitalization(hospitalization model.Hospitalization) error {
	if err := ValidateHospitalization(hospitalization); err != nil {
		return err
	}

	hospitalization.CreatedAt = time.Now()
	hospitalization.UpdatedAt = time.Now()

	repo := gethospitalizationRepo()
	return repo.SaveHospitalization(&hospitalization)
}

