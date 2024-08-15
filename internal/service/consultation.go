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

func AddConsultationTransaction(consultation model.Consultation, sender, receiver string, amount float64) error {

	// Validação da consulta
	if err := ValidateConsultation(consultation); err != nil {
		log.Printf("Erro ao validar consulta: %v", err)
		return err
	}

	// Converta o objeto Consultation para JSON
	consultationJSON, err := json.Marshal(consultation)
	if err != nil {
		log.Printf("Erro ao converter consulta para JSON: %v", err)
		return err // Retorna o erro se a conversão falhar
	}

	// Crie uma nova transação para a consulta com o JSON
	transaction := blockchain.NewTransaction(sender, receiver, amount, string(consultationJSON))

	log.Printf("Adicionando transação: %v", transaction)

	// newBlock := Block{
	// 	Index:        len(Blockchain) + 1,
	// 	Timestamp:    time.Now(),
	// 	Transactions: []Transaction{transaction},
	// 	PreviousHash: Blockchain[len(Blockchain)-1].Hash,
	// }
	newBlock := blockchain.NewBlock(len(blockchain.Blockchain)+1, []blockchain.Transaction{*transaction}, blockchain.Blockchain[len(blockchain.Blockchain)-1].Hash)

	log.Printf("[%v] Novo bloco criado: %v", time.Now().Format(time.RFC3339), newBlock)

	// Minerar o bloco e adicioná-lo à blockchain
	difficulty := 2
	newBlock.MineBlock(difficulty)
	blockchain.Blockchain = append(blockchain.Blockchain, *newBlock)

	log.Printf("Bloco adicionado à blockchain: %v", newBlock)

	return nil
}

// Função para buscar uma consulta por ID na blockchain
func GetConsultationByID(id uuid.UUID) (*model.Consultation, error) {
	log.Printf("Buscando consulta por ID: %v", id)
	// Itera sobre cada bloco na blockchain
	for _, block := range blockchain.Blockchain {
		for _, transaction := range block.Transactions {
			var consultation model.Consultation
			err := json.Unmarshal([]byte(transaction.Data), &consultation)
			if err != nil {
				log.Printf("Erro ao decodificar transação: %v", err)
				return nil, err
			}
			if consultation.ID == id {
				log.Printf("Consulta encontrada: %v", consultation)
				return &consultation, nil
			}
		}
	}
	log.Printf("Consulta não encontrada para o ID: %v", id)
	return nil, nil
}

// Função para buscar uma consulta por ID do animal na blockchain
func GetConsultationByAnimalID(id uuid.UUID) ([]model.Consultation, error) {
	log.Printf("Buscando consultas por Animal ID: %v", id)
	var consultations []model.Consultation
	for _, block := range blockchain.Blockchain {
		for _, transaction := range block.Transactions {
			var consultation model.Consultation
			err := json.Unmarshal([]byte(transaction.Data), &consultation)
			if err != nil {
				log.Printf("Erro ao decodificar transação: %v", err)
				return nil, err
			}
			if consultation.AnimalID == id {
				consultations = append(consultations, consultation)
				log.Printf("Consulta encontrada: %v", consultation)
			}
		}
	}
	log.Printf("Total de consultas encontradas para Animal ID %v: %d", id, len(consultations))
	return consultations, nil
}

// Função para buscar uma consulta por ID do veterinário na blockchain
func GetConsultationByVeterinaryCRVM(crvm int) ([]model.Consultation, error) {
	log.Printf("Buscando consultas por Veterinary ID: %v", crvm)
	var consultations []model.Consultation
	for _, block := range blockchain.Blockchain {
		for _, transaction := range block.Transactions {
			var consultation model.Consultation
			err := json.Unmarshal([]byte(transaction.Data), &consultation)
			if err != nil {
				log.Printf("Erro ao decodificar transação: %v", err)
				return nil, err
			}
			if consultation.CRVM == crvm{
				consultations = append(consultations, consultation)
				log.Printf("Consulta encontrada: %v", consultation)
			}
		}
	}
	log.Printf("Total de consultas encontradas para Veterinary ID %v: %d", crvm, len(consultations))
	return consultations, nil
}

// Função para agendar consulta
func ScheduleConsultation(consultation model.Consultation, sender, receiver string, amount float64) error {
	log.Printf("Agendando consulta: %v", consultation)
	consultation.ConsultationStatus = "Scheduled"
	return AddConsultationTransaction(consultation, sender, receiver, amount)
}

// Função para cancelar consulta
func CancelConsultation(consultation model.Consultation, sender, receiver string, amount float64) error {
	log.Printf("Cancelando consulta: %v", consultation)
	consultation.ConsultationStatus = "Canceled"
	return AddConsultationTransaction(consultation, sender, receiver, amount)
}

// Função para confirmar consulta
func ConfirmConsultation(consultation model.Consultation, sender, receiver string, amount float64) error {
	log.Printf("Confirmando consulta: %v", consultation)
	consultation.ConsultationStatus = "Confirmed"
	return AddConsultationTransaction(consultation, sender, receiver, amount)
}

// Função para atualizar consulta
func UpdateConsultation(consultation model.Consultation, sender, receiver string, amount float64) error {
	log.Printf("Atualizando consulta: %v", consultation)
	return AddConsultationTransaction(consultation, sender, receiver, amount)
}

func AddConsultationHistory(consultation model.Consultation, changes []model.Change) {
	history := model.ConsultationHistory{
		ConsultationID: consultation.ID,
		Changes:        changes,
		Timestamp:      time.Now(),
	}
	// Aqui, você pode armazenar o histórico na blockchain ou em um armazenamento separado
	log.Printf("Histórico de consulta adicionado: %v", history)
}

func TrackChanges(oldConsultation, newConsultation model.Consultation) []model.Change {
	var changes []model.Change
	// Comparar os campos relevantes e adicionar ao slice de mudanças
	if oldConsultation.ConsultationStatus != newConsultation.ConsultationStatus {
		changes = append(changes, model.Change{
			Field:    "ConsultationStatus",
			OldValue: oldConsultation.ConsultationStatus,
			NewValue: newConsultation.ConsultationStatus,
		})
	}
	// Adicionar mais comparações conforme necessário
	return changes
}

func ValidateConsultation(consultation model.Consultation) error {
	if consultation.AnimalID == [16]byte{} {
		return errors.New("AnimalID não pode ser zero")
	}
	if consultation.CRVM == 0 {
		return errors.New("VeterinaryID não pode ser zero")
	}
	if consultation.ConsultationPrice < 0 {
		return errors.New("ConsultationPrice não pode ser negativo")
	}
	// Adicione outras validações conforme necessário
	return nil
}
