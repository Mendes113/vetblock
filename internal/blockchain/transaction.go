package blockchain

import "time"

type Transaction struct {
    Sender    string    `json:"sender"`
    Receiver  string    `json:"receiver"`
    Amount    float64   `json:"amount"`
    Timestamp time.Time `json:"timestamp"`
    Data      string    `json:"data"`
}

type AnimalTransaction struct {
    Animal Animal `json:"animal"`
    Timestamp time.Time `json:"timestamp"`
}


// // Estrutura para transações de medicações
// type MedicationTransaction struct {
//     Medication Medication `json:"medication"`
//     Timestamp time.Time `json:"timestamp"`
// }