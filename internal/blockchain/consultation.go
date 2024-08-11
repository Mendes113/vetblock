package blockchain

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type Consultation struct {
	ID                      uuid.UUID `json:"id"`
	AnimalID                uint64    `json:"animal_id"`
	VeterinaryID            uint64    `json:"veterinary_id"`
	ConsultationDate        string    `json:"consultation_date"`
	ConsultationHour        string    `json:"consultation_hour"`
	ConsultationType        string    `json:"consultation_type"`
	ConsultationDescription string    `json:"consultation_description"`
	ConsultationPrescription string   `json:"consultation_prescription"`
	ConsultationPrice       float64   `json:"consultation_price"`
	ConsultationStatus      string    `json:"consultation_status"`
}


type ConsultationHistory struct {
    ConsultationID uuid.UUID `json:"consultation_id"`
    Changes        []Change  `json:"changes"`
    Timestamp      time.Time `json:"timestamp"`
}

type Change struct {
    Field    string `json:"field"`
    OldValue string `json:"old_value"`
    NewValue string `json:"new_value"`
}




func AddConsultationTransaction(consultation Consultation, sender, receiver string, amount float64) error {

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
	transaction := Transaction{
		Sender:    sender,
		Receiver:  receiver,
		Amount:    amount,
		Timestamp: time.Now(),
		Data:      string(consultationJSON), // Armazena o JSON como uma string
	}

	log.Printf("Adicionando transação: %v", transaction)

	newBlock := Block{
		Index:        len(Blockchain) + 1,
		Timestamp:    time.Now(),
		Transactions: []Transaction{transaction},
		PreviousHash: Blockchain[len(Blockchain)-1].Hash,
	}

	log.Printf("[%v] Novo bloco criado: %v", time.Now().Format(time.RFC3339), newBlock)

	// Minerar o bloco e adicioná-lo à blockchain
	difficulty := 2
	newBlock.MineBlock(difficulty)
	Blockchain = append(Blockchain, newBlock)

	log.Printf("Bloco adicionado à blockchain: %v", newBlock)

	return nil
}

// Função para buscar uma consulta por ID na blockchain
func GetConsultationByID(id uuid.UUID) (*Consultation, error) {
	log.Printf("Buscando consulta por ID: %v", id)
	// Itera sobre cada bloco na blockchain
	for _, block := range Blockchain {
		for _, transaction := range block.Transactions {
			var consultation Consultation
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
func GetConsultationByAnimalID(id uint64) ([]Consultation, error) {
	log.Printf("Buscando consultas por Animal ID: %v", id)
	var consultations []Consultation
	for _, block := range Blockchain {
		for _, transaction := range block.Transactions {
			var consultation Consultation
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
func GetConsultationByVeterinaryID(id uint64) ([]Consultation, error) {
	log.Printf("Buscando consultas por Veterinary ID: %v", id)
	var consultations []Consultation
	for _, block := range Blockchain {
		for _, transaction := range block.Transactions {
			var consultation Consultation
			err := json.Unmarshal([]byte(transaction.Data), &consultation)
			if err != nil {
				log.Printf("Erro ao decodificar transação: %v", err)
				return nil, err
			}
			if consultation.VeterinaryID == id {
				consultations = append(consultations, consultation)
				log.Printf("Consulta encontrada: %v", consultation)
			}
		}
	}
	log.Printf("Total de consultas encontradas para Veterinary ID %v: %d", id, len(consultations))
	return consultations, nil
}

// Função para agendar consulta
func ScheduleConsultation(consultation Consultation, sender, receiver string, amount float64) error {
	log.Printf("Agendando consulta: %v", consultation)
	consultation.ConsultationStatus = "Scheduled"
	return AddConsultationTransaction(consultation, sender, receiver, amount)
}

// Função para cancelar consulta
func CancelConsultation(consultation Consultation, sender, receiver string, amount float64) error {
	log.Printf("Cancelando consulta: %v", consultation)
	consultation.ConsultationStatus = "Canceled"
	return AddConsultationTransaction(consultation, sender, receiver, amount)
}

// Função para confirmar consulta
func ConfirmConsultation(consultation Consultation, sender, receiver string, amount float64) error {
	log.Printf("Confirmando consulta: %v", consultation)
	consultation.ConsultationStatus = "Confirmed"
	return AddConsultationTransaction(consultation, sender, receiver, amount)
}

// Função para atualizar consulta
func UpdateConsultation(consultation Consultation, sender, receiver string, amount float64) error {
	log.Printf("Atualizando consulta: %v", consultation)
	return AddConsultationTransaction(consultation, sender, receiver, amount)
}

func AddConsultationHistory(consultation Consultation, changes []Change) {
    history := ConsultationHistory{
        ConsultationID: consultation.ID,
        Changes:        changes,
        Timestamp:      time.Now(),
    }
    // Aqui, você pode armazenar o histórico na blockchain ou em um armazenamento separado
    log.Printf("Histórico de consulta adicionado: %v", history)
}

func TrackChanges(oldConsultation, newConsultation Consultation) []Change {
    var changes []Change
    // Comparar os campos relevantes e adicionar ao slice de mudanças
    if oldConsultation.ConsultationStatus != newConsultation.ConsultationStatus {
        changes = append(changes, Change{
            Field:    "ConsultationStatus",
            OldValue: oldConsultation.ConsultationStatus,
            NewValue: newConsultation.ConsultationStatus,
        })
    }
    // Adicionar mais comparações conforme necessário
    return changes
}


func ValidateConsultation(consultation Consultation) error {
    if consultation.AnimalID == 0 {
        return errors.New("AnimalID não pode ser zero")
    }
    if consultation.VeterinaryID == 0 {
        return errors.New("VeterinaryID não pode ser zero")
    }
    if consultation.ConsultationPrice < 0 {
        return errors.New("ConsultationPrice não pode ser negativo")
    }
    // Adicione outras validações conforme necessário
    return nil
}
