package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Animal struct {
	ID          uuid.UUID       `gorm:"type:uuid;primary_key;" json:"animal_id"`
	Name        string          `json:"name" gorm:"not null" validate:"required,min=2,max=100"`
	Species     string          `json:"species" gorm:"not null" validate:"required"`
	Breed       string          `json:"breed" gorm:"not null" validate:"required"`
	Age         int             `json:"age" validate:"gte=0"`
	Description string          `json:"description"`
	Timestamp   time.Time       `json:"timestamp" gorm:"autoCreateTime"`
	CreatedAt   time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt  `gorm:"index" json:"-"` // Soft delete
	CPFTutor    string          `gorm:"type:char(11);not null" json:"cpf_tutor" validate:"required,len=11"`
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
    ID                      uuid.UUID      `gorm:"type:uuid;primary_key" json:"consultation_id"`
    AnimalID                uuid.UUID      `gorm:"type:uuid;not null" json:"animal_id" validate:"required,uuid"`
    CRVM                    string            `gorm:"column:crvm;not null" json:"crvm" validate:"required,min=1"`
    ConsultationDate        time.Time      `json:"consultation_date" validate:"required"`
    ConsultationHour        string         `json:"consultation_hour" validate:"required,len=5,datetime=15:04"` // Ajuste o formato se necessário
    Observation             string         `json:"observation" validate:"max=255"`
    Reason                  string         `json:"reason" validate:"required,min=10,max=255"`
    ConsultationType        string         `json:"consultation_type" validate:"required"`
    ConsultationDescription string         `json:"consultation_description" validate:"required"`
    ConsultationPrescription string        `json:"consultation_prescription"`
    ConsultationPrice       float64        `json:"consultation_price" validate:"required,gte=0"`
    ConsultationStatus      string         `json:"consultation_status" validate:"required,oneof=scheduled completed canceled"`
    CreatedAt               time.Time      `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt               time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
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
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;" json:"hospitalization_id"`
	PatientID   uuid.UUID      `gorm:"type:uuid;not null" json:"patient_id" validate:"required,uuid"`
	StartDate   time.Time      `json:"start_date" validate:"required"`
	EndDate     time.Time      `json:"end_date" validate:"required,gtfield=StartDate"` // Valida que EndDate é depois de StartDate
	Reason      string         `json:"reason" validate:"required,min=10,max=255"`
	CRVM        int            `json:"doctor_id" validate:"required,min=1"`
	Medications []string       `gorm:"type:jsonb" json:"medications" validate:"dive,required,min=1"` // Valida que cada medicação está presente
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}

