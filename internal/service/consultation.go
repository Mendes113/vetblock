package service

import (
	
	"vetblock/internal/db/model"

	
)

func AddConsultation(consultation model.Consultation, sender, receiver string, amount float64) error {

	// // Validação da consulta
	// if err := ValidateConsultation(consultation); err != nil {
	// 	log.Printf("Erro ao validar consulta: %v", err)
	// 	return err
	// }

	// // Converta o objeto Consultation para JSON
	// consultationJSON, err := json.Marshal(consultation)
	// if err != nil {
	// 	log.Printf("Erro ao converter consulta para JSON: %v", err)
	// 	return err // Retorna o erro se a conversão falhar
	// }


	

	return nil
}


// // Função para buscar uma consulta por ID do animal na blockchain
// func GetConsultationByAnimalID(id uuid.UUID) ([]model.Consultation, error) {
// 	log.Printf("Buscando consultas por Animal ID: %v", id)
// 	var consultations []model.Consultation
// 	for _, block := range blockchain.Blockchain {
// 		for _, transaction := range block.Transactions {
// 			var consultation model.Consultation
// 			err := json.Unmarshal([]byte(transaction.Data), &consultation)
// 			if err != nil {
// 				log.Printf("Erro ao decodificar transação: %v", err)
// 				return nil, err
// 			}
// 			if consultation.AnimalID == id {
// 				consultations = append(consultations, consultation)
// 				log.Printf("Consulta encontrada: %v", consultation)
// 			}
// 		}
// 	}
// 	log.Printf("Total de consultas encontradas para Animal ID %v: %d", id, len(consultations))
// 	return consultations, nil
// }

// // Função para buscar uma consulta por ID do veterinário na blockchain
// func GetConsultationByVeterinaryCRVM(crvm int) ([]model.Consultation, error) {
// 	log.Printf("Buscando consultas por Veterinary ID: %v", crvm)
// 	var consultations []model.Consultation
// 	for _, block := range blockchain.Blockchain {
// 		for _, transaction := range block.Transactions {
// 			var consultation model.Consultation
// 			err := json.Unmarshal([]byte(transaction.Data), &consultation)
// 			if err != nil {
// 				log.Printf("Erro ao decodificar transação: %v", err)
// 				return nil, err
// 			}
// 			if consultation.CRVM == crvm{
// 				consultations = append(consultations, consultation)
// 				log.Printf("Consulta encontrada: %v", consultation)
// 			}
// 		}
// 	}
// 	log.Printf("Total de consultas encontradas para Veterinary ID %v: %d", crvm, len(consultations))
// 	return consultations, nil
// }

// // Função para agendar consulta
// func ScheduleConsultation(consultation model.Consultation, sender, receiver string, amount float64) error {
// 	log.Printf("Agendando consulta: %v", consultation)
// 	consultation.ConsultationStatus = "Scheduled"
// 	return AddConsultationTransaction(consultation, sender, receiver, amount)
// }

// // Função para cancelar consulta
// func CancelConsultation(consultation model.Consultation, sender, receiver string, amount float64) error {
// 	log.Printf("Cancelando consulta: %v", consultation)
// 	consultation.ConsultationStatus = "Canceled"
// 	return AddConsultationTransaction(consultation, sender, receiver, amount)
// }

// // Função para confirmar consulta
// func ConfirmConsultation(consultation model.Consultation, sender, receiver string, amount float64) error {
// 	log.Printf("Confirmando consulta: %v", consultation)
// 	consultation.ConsultationStatus = "Confirmed"
// 	return AddConsultationTransaction(consultation, sender, receiver, amount)
// }

// // Função para atualizar consulta
// func UpdateConsultation(consultation model.Consultation, sender, receiver string, amount float64) error {
// 	log.Printf("Atualizando consulta: %v", consultation)
// 	return AddConsultationTransaction(consultation, sender, receiver, amount)
// }

// func AddConsultationHistory(consultation model.Consultation, changes []model.Change) {
// 	history := model.ConsultationHistory{
// 		ConsultationID: consultation.ID,
// 		Changes:        changes,
// 		Timestamp:      time.Now(),
// 	}
// 	// Aqui, você pode armazenar o histórico na blockchain ou em um armazenamento separado
// 	log.Printf("Histórico de consulta adicionado: %v", history)
// }

// func TrackChanges(oldConsultation, newConsultation model.Consultation) []model.Change {
// 	var changes []model.Change
// 	// Comparar os campos relevantes e adicionar ao slice de mudanças
// 	if oldConsultation.ConsultationStatus != newConsultation.ConsultationStatus {
// 		changes = append(changes, model.Change{
// 			Field:    "ConsultationStatus",
// 			OldValue: oldConsultation.ConsultationStatus,
// 			NewValue: newConsultation.ConsultationStatus,
// 		})
// 	}
// 	// Adicionar mais comparações conforme necessário
// 	return changes
// }

// func ValidateConsultation(consultation model.Consultation) error {
// 	if consultation.AnimalID == [16]byte{} {
// 		return errors.New("AnimalID não pode ser zero")
// 	}
// 	if consultation.CRVM == 0 {
// 		return errors.New("VeterinaryID não pode ser zero")
// 	}
// 	if consultation.ConsultationPrice < 0 {
// 		return errors.New("ConsultationPrice não pode ser negativo")
// 	}
// 	// Adicione outras validações conforme necessário
// 	return nil
// }
