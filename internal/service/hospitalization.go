package service

import (
	"encoding/json"
	"errors"
	"log"
	"time"
	"vetblock/internal/blockchain"
	"vetblock/internal/db/model"

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

func ValidateHospitalization(hospitalization model.Hospitalization) error {
	validations := []struct {
		condition bool
		errMsg    string
	}{
		{hospitalization.ID == uuid.UUID{}, ErrInvalidHospitalizationID},
		{hospitalization.PatientID == uuid.UUID{}, ErrInvalidPatientID},
		{hospitalization.StartDate == time.Time{}, ErrInvalidStartDate},
		{hospitalization.EndDate == time.Time{}, ErrInvalidEndDate},
		{hospitalization.Reason == "", ErrInvalidReason},
		{hospitalization.DoctorID == uuid.UUID{}, ErrInvalidDoctorID},
		{len(hospitalization.Medications) == 0, ErrInvalidMedications},
	}

	for _, v := range validations {
		if v.condition {
			return errors.New(v.errMsg)
		}
	}

	return nil
}

func AddHospitalizationTransaction(hospitalization model.Hospitalization, sender, receiver string, amount float64) error {
	// Validação da hospitalização
	if err := ValidateHospitalization(hospitalization); err != nil {
		log.Printf("Erro ao validar hospitalização: %v", err)
		return err
	}

	hospitalizationJSON, err := json.Marshal(hospitalization)
	if err != nil {
		log.Printf("Erro ao converter hospitalização para JSON: %v", err)
		return err
	}

	transaction := blockchain.NewTransaction(sender, receiver, amount, string(hospitalizationJSON))

	// newBlock := Block{
	// 	Index:        len(Blockchain) + 1,
	// 	Timestamp:    time.Now(),
	// 	Transactions: []Transaction{transaction},
	// 	Hash:         "",
	// 	PreviousHash: Blockchain[len(Blockchain)-1].Hash,
	// }

	newBlock := blockchain.NewBlock(len(blockchain.Blockchain)+1, []blockchain.Transaction{*transaction}, blockchain.Blockchain[len(blockchain.Blockchain)-1].Hash)

	difficulty := 2
	newBlock.MineBlock(difficulty)
	blockchain.Blockchain = append(blockchain.Blockchain, *newBlock)
	log.Printf("Bloco adicionado à blockchain: %v", newBlock)

	return nil
}

func GetHospitalizationByID(id uuid.UUID) (*model.Hospitalization, error) {
	log.Printf("Buscando hospitalização por ID: %v", id)
	for _, block := range blockchain.Blockchain {
		for _, transaction := range block.Transactions {
			var hospitalization model.Hospitalization
			err := json.Unmarshal([]byte(transaction.Data), &hospitalization)
			if err != nil {
				log.Printf("Erro ao decodificar hospitalização: %v", err)
				continue
			}
			if hospitalization.ID == id {
				return &hospitalization, nil
			}
		}
	}
	return nil, errors.New("hospitalization not found")
}

func GetHospitalizationsByPatientID(patientID uuid.UUID) ([]model.Hospitalization, error) {
	log.Printf("Buscando hospitalizações por ID do paciente: %v", patientID)
	var hospitalizations []model.Hospitalization
	for _, block := range blockchain.Blockchain {
		for _, transaction := range block.Transactions {
			var hospitalization model.Hospitalization
			err := json.Unmarshal([]byte(transaction.Data), &hospitalization)
			if err != nil {
				log.Printf("Erro ao decodificar hospitalização: %v", err)
				continue
			}
			if hospitalization.PatientID == patientID {
				hospitalizations = append(hospitalizations, hospitalization)
			}
		}
	}
	return hospitalizations, nil
}

func GetHospitalizationsByDoctorID(doctorID uuid.UUID) ([]model.Hospitalization, error) {
	log.Printf("Buscando hospitalizações por ID do médico: %v", doctorID)
	var hospitalizations []model.Hospitalization
	for _, block := range blockchain.Blockchain {
		for _, transaction := range block.Transactions {
			var hospitalization model.Hospitalization
			err := json.Unmarshal([]byte(transaction.Data), &hospitalization)
			if err != nil {
				log.Printf("Erro ao decodificar hospitalização: %v", err)
				continue
			}
			if hospitalization.DoctorID == doctorID {
				hospitalizations = append(hospitalizations, hospitalization)
			}
		}
	}
	return hospitalizations, nil
}

// talvez deva ir para medication.go
func GetMedicationByHospitalizationID(id uuid.UUID) ([]string, error) {
	log.Printf("Buscando medicamentos por ID da hospitalização: %v", id)
	for _, block := range blockchain.Blockchain {
		for _, transaction := range block.Transactions {
			var hospitalization model.Hospitalization
			err := json.Unmarshal([]byte(transaction.Data), &hospitalization)
			if err != nil {
				log.Printf("Erro ao decodificar hospitalização: %v", err)
				continue
			}
			if hospitalization.ID == id {
				return hospitalization.Medications, nil
			}
		}
	}
	return nil, errors.New("medication not found")
}

func GetHospitalizations() ([]model.Hospitalization, error) {
	log.Printf("Buscando todas as hospitalizações")
	var hospitalizations []model.Hospitalization
	for _, block := range blockchain.Blockchain {
		for _, transaction := range block.Transactions {
			var hospitalization model.Hospitalization
			err := json.Unmarshal([]byte(transaction.Data), &hospitalization)
			if err != nil {
				log.Printf("Erro ao decodificar hospitalização: %v", err)
				continue
			}
			hospitalizations = append(hospitalizations, hospitalization)
		}
	}
	return hospitalizations, nil
}
