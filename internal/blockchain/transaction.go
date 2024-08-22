package blockchain

import (
	"fmt"
	"time"
	"vetblock/internal/db/model"

	"github.com/google/uuid"
)

type Transaction struct {
	Sender    string    `json:"sender"`
	Receiver  string    `json:"receiver"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
}

type AnimalTransaction struct {
	Animal    model.Animal `json:"animal"`
	Timestamp time.Time    `json:"timestamp"`
}

func (t *Transaction) String() string {
	return t.Sender + t.Receiver + t.Timestamp.String() + t.Data
}

// create a new transaction
func NewTransaction(sender, receiver string, data string) *Transaction {
	return &Transaction{
		Sender:    sender,
		Receiver:  receiver,
		Timestamp: time.Now(),
		Data:      data,
	}
}

// Valida uma transação de dados
func (t *Transaction) Validate() bool {
    // Permitir uma margem de 5 minutos para o timestamp
    if t.Timestamp.After(time.Now().Add(5 * time.Minute)) {
        fmt.Println("Transação inválida: timestamp está no futuro")
        return false
    }

    if t.Sender == t.Receiver {
        fmt.Println("Transação inválida: remetente e receptor são iguais")
        return false
    }

    if len(t.Data) == 0 {
        fmt.Println("Transação inválida: dados estão vazios")
        return false
    }

    return true
}


// Valida uma transação de animal
func (at *AnimalTransaction) Validate() bool {
	// Verificar se a transação de animal possui uma referência válida
	if at.Animal.ID == uuid.Nil {
		fmt.Println("Transação de animal inválida: ID do animal está vazio")
		return false
	}

	// Verificar se o timestamp não está no futuro
	if at.Timestamp.After(time.Now()) {
		fmt.Println("Transação de animal inválida: timestamp está no futuro")
		return false
	}

	

	return true
}
