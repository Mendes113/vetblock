package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Animal struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;" json:"animal_id"`
	Name        string         `json:"name" gorm:"not null" validate:"required,min=2,max=100"`
	Species     string         `json:"species" gorm:"not null" validate:"required"`
	Breed       string         `json:"breed" gorm:"not null" validate:"required"`
	Age         int            `json:"age" validate:"gte=0"`
	Weight      float64        `json:"weight" validate:"gte=0"`
	Description string         `json:"description"`
	Timestamp   time.Time      `json:"timestamp" gorm:"autoCreateTime"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
	CPFTutor    string         `gorm:"type:char(11);not null" json:"cpf_tutor" validate:"required,len=11"`
}

type ConsultationDosage struct {
    ConsultationID uuid.UUID `gorm:"type:uuid;primary_key;" json:"consultation_id"`
    DosageID       uuid.UUID `gorm:"type:uuid;primary_key;" json:"dosage_id"`
}

type HospitalizationDosage struct {
    HospitalizationID uuid.UUID `gorm:"type:uuid;primary_key;" json:"hospitalization_id"`
    DosageID          uuid.UUID `gorm:"type:uuid;primary_key;" json:"dosage_id"`
}

type Dosage struct {
    ID                 uuid.UUID      `gorm:"type:uuid;primary_key;" json:"dosage_id"`
    AnimalID           uuid.UUID      `gorm:"type:uuid;not null" json:"animal_id" validate:"required,uuid"`
    MedicationID       uuid.UUID      `gorm:"type:uuid;not null" json:"medication_id" validate:"required,uuid"`
    StartDate          time.Time      `json:"start_date" validate:"required"`
    EndDate            time.Time      `json:"end_date" validate:"required,gtfield=StartDate"`
    Quantity           int            `json:"quantity" validate:"gte=0"`
    Dosage             string         `json:"dosage" validate:"required"`
    ConsultationID     *uuid.UUID     `gorm:"type:uuid" json:"consultation_id"` // Relacionamento opcional
    HospitalizationID  *uuid.UUID     `gorm:"type:uuid" json:"hospitalization_id"` // Relacionamento opcional
    CreatedAt          time.Time      `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt          time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
    DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
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
	ID                       uuid.UUID      `gorm:"type:uuid;primary_key" json:"consultation_id"`
	AnimalID                 uuid.UUID      `gorm:"type:uuid;not null" json:"animal_id" validate:"required,uuid"`
	CRVM                     string         `gorm:"column:crvm;not null" json:"crvm" validate:"required,min=1"`
	ConsultationDate         time.Time      `json:"consultation_date" validate:"required"`
	ConsultationHour         string         `json:"consultation_hour" validate:"required,len=5,datetime=15:04"` // Ajuste o formato se necessário
	Observation              string         `json:"observation" validate:"max=255"`
	Reason                   string         `json:"reason" validate:"required,min=10,max=255"`
	ConsultationType         string         `json:"consultation_type" validate:"required"`
	ConsultationDescription  string         `json:"consultation_description" validate:"required"`
	ConsultationPrescription string         `json:"consultation_prescription"`
	ConsultationPrice        float64        `json:"consultation_price" validate:"required,gte=0"`
	ConsultationStatus       string         `json:"consultation_status" validate:"required,oneof=scheduled completed canceled"`
	CreatedAt                time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt                time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt                gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}

type ConsultationHistory struct {
	ID             uuid.UUID      `gorm:"type:uuid;primary_key;" json:"consultation_history_id"`
	ConsultationID uuid.UUID      `gorm:"type:uuid;not null" json:"consultation_id"`
	Changes        []Change       `gorm:"type:jsonb" json:"changes"` // Use JSONB for arrays
	Timestamp      time.Time      `json:"timestamp"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
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

type Medication struct {
	ID                   uuid.UUID      `gorm:"type:uuid;primary_key;" json:"medication_id"` // Identificador único do medicamento, usando UUID como chave primária.
	Name                 string         `json:"name" gorm:"not null" validate:"required,min=2,max=100"` // Nome do medicamento, obrigatório, com tamanho mínimo de 2 e máximo de 100 caracteres.
	Description          string         `json:"description" validate:"max=255"` // Descrição do medicamento, opcional, com tamanho máximo de 255 caracteres.
	Price                float64        `json:"price" validate:"required,gte=0"` // Preço do medicamento, obrigatório, deve ser maior ou igual a zero.
	BatchNumber          string         `json:"batch_number" validate:"required"` // Número do lote do medicamento, obrigatório.
	Concentration        string         `json:"concentration" validate:"required, min=2"` // Concentração do medicamento, obrigatório, com tamanho mínimo de 2 caracteres.
	Presentation         string         `json:"presentation" validate:"required"` // Forma de apresentação do medicamento (ex: comprimidos, líquido), obrigatório.
	DosageForm           string         `json:"dosage_form" validate:"required"` // Forma de dosagem do medicamento (ex: oral, injetável), obrigatório.
	ActivePrinciples     pq.StringArray `json:"active_principles" gorm:"type:text[]" validate:"required"` // Lista dos princípios ativos do medicamento, obrigatório.
	Manufacturer         string         `json:"manufacturer" validate:"required"` // Nome do fabricante do medicamento, obrigatório.
	Quantity             int            `json:"quantity" validate:"gte=0"` // Quantidade em estoque do medicamento, deve ser maior ou igual a zero.
	Unit                 string         `json:"unit" validate:"required"` // Unidade de medida do medicamento (ex: mg, ml), obrigatório.
	StorageConditions    string         `json:"storage_conditions"` // Condições de armazenamento do medicamento, opcional.
	PrescriptionRequired bool           `json:"prescription_required"` // Indica se o medicamento requer prescrição médica, booleano.
	Expiration           time.Time      `json:"expiration" validate:"required"` // Data de validade do medicamento, obrigatório.
	CreatedAt            time.Time      `json:"created_at" gorm:"autoCreateTime"` // Data de criação do registro, automaticamente preenchido pelo GORM.
	UpdatedAt            time.Time      `json:"updated_at" gorm:"autoUpdateTime"` // Data de atualização do registro, automaticamente preenchido pelo GORM.
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"` // Campo usado para soft delete, permitindo que o registro seja marcado como deletado sem ser removido fisicamente do banco.
}

