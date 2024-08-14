package blockchain

import (
	"strconv"
	"time"
	"vetblock/internal/db"
)

type Transaction struct {
    Sender    string    `json:"sender"`
    Receiver  string    `json:"receiver"`
    Amount    float64   `json:"amount"`
    Timestamp time.Time `json:"timestamp"`
    Data      string    `json:"data"`
}

type AnimalTransaction struct {
    Animal db.Animal `json:"animal"`
    Timestamp time.Time `json:"timestamp"`
}



func (t *Transaction) String() string {
    return t.Sender + t.Receiver + strconv.FormatFloat(t.Amount, 'f', -1, 64) + t.Timestamp.String() + t.Data
}

//create a new transaction
func NewTransaction(sender, receiver string, amount float64, data string) *Transaction {
    return &Transaction{
        Sender:    sender,
        Receiver:  receiver,
        Amount:    amount,
        Timestamp: time.Now(),
        Data:      data,
    }
}



// // Estrutura para transações de medicações
// type MedicationTransaction struct {
//     Medication Medication `json:"medication"`
//     Timestamp time.Time `json:"timestamp"`
// }