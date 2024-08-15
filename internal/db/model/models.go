package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Animal struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;" json:"animal_id"`
	Name        string    `json:"name"`
	Species     string    `json:"species"`
	Breed       string    `json:"breed"`
	Age         int       `json:"age"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
	CPFTutor     uuid.UUID `gorm:"type:uuid;not null" json:"CPFtutor"`
}


type CustomDate time.Time

func (cd *CustomDate) UnmarshalJSON(data []byte) error {
    str := strings.Trim(string(data), `"`)
    t, err := time.Parse("2006-01-02", str)
    if err != nil {
        return err
    }
    *cd = CustomDate(t)
    return nil
}

func (cd CustomDate) MarshalJSON() ([]byte, error) {
    return []byte(`"` + time.Time(cd).Format("2006-01-02") + `"`), nil
}



type Consultation struct {
    ID                      uuid.UUID      `gorm:"type:uuid;primary_key;" json:"consultation_id"`
    AnimalID                uuid.UUID      `gorm:"type:uuid;not null" json:"animal_id"`
    CRVM           			int      `gorm:"type:uuid;not null" json:"crvm"`
    ConsultationDate        CustomDate     `json:"consultation_date"`
    ConsultationHour        string         `json:"consultation_hour"`
    ConsultationType        string         `json:"consultation_type"`
    ConsultationDescription string         `json:"consultation_description"`
    ConsultationPrescription string        `json:"consultation_prescription"`
    ConsultationPrice       float64        `json:"consultation_price"`
    ConsultationStatus      string         `json:"consultation_status"`
    CreatedAt               time.Time      `json:"created_at"`
    UpdatedAt               time.Time      `json:"updated_at"`
    DeletedAt               gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}

type ConsultationHistory struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;" json:"consultation_history_id"`
	ConsultationID uuid.UUID `gorm:"type:uuid;not null" json:"consultation_id"`
	Changes        []Change  `gorm:"type:jsonb" json:"changes"` // Use JSONB for arrays
	Timestamp      time.Time `json:"timestamp"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}

type Change struct {
	Field    string `json:"field"`
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}

type Hospitalization struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;" json:"hospitalization_id"`
	PatientID   uuid.UUID `gorm:"type:uuid;not null" json:"patient_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Reason      string    `json:"reason"`
	CRVM    	 int `gorm:"type:uuid;not null" json:"doctor_id"`
	Medications []string  `gorm:"type:jsonb" json:"medications"` // Use JSONB for arrays
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}


