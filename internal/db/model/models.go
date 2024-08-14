package model

import (
	"time"

	"github.com/google/uuid"
)


type Animal struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Species     string    `json:"species"`
    Breed       string    `json:"breed"`
    Age         int       `json:"age"`
    Description string    `json:"description"`
    Timestamp   time.Time `json:"timestamp"`
}


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


type Hospitalization struct {
	ID          string   `json:"id"`
	PatientID   string   `json:"patient_id"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Reason      string   `json:"reason"`
	DoctorID    string   `json:"doctor_id"`
	Medications []string `json:"medications"`
}


