package model

import (
	"errors"
	"regexp"
	"time"

	"gorm.io/gorm"
)

type Veterinary struct {
	CRVM      string         `gorm:"type:char(12);primary_key" json:"crvm"`
	Name      string         `json:"name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete
}

func isValidCRVM(crvm string) bool {
	re := regexp.MustCompile(`^[0-9]{6,8}-[A-Z]{2}$`)
	return re.MatchString(crvm)
}

func NewVeterinary(crvm, name, lastName, email, phone string) (*Veterinary, error) {
	if !isValidCRVM(crvm) {
		return nil, errors.New("CRVM inválido")
	}
	return &Veterinary{
		CRVM:      crvm,
		Name:      name,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (v *Veterinary) GetCRVM() string {
	return v.CRVM
}

func (v *Veterinary) GetName() string {
	return v.Name
}

func (v *Veterinary) GetLastName() string {
	return v.LastName
}

func (v *Veterinary) GetEmail() string {
	return v.Email
}

func (v *Veterinary) GetPhone() string {
	return v.Phone
}

func (v *Veterinary) SetCRVM(crvm string) error {
	if !isValidCRVM(crvm) {
		return errors.New("CRVM inválido")
	}
	v.CRVM = crvm
	return nil
}

func (v *Veterinary) SetName(name string) {
	v.Name = name
}

func (v *Veterinary) SetLastName(lastName string) {
	v.LastName = lastName
}

func (v *Veterinary) SetEmail(email string) {
	v.Email = email
}

func (v *Veterinary) SetPhone(phone string) {
	v.Phone = phone
}

func (v *Veterinary) Update(name, lastName, email, phone string) {
	v.Name = name
	v.LastName = lastName
	v.Email = email
	v.Phone = phone
}

func (v *Veterinary) Validate() error {
	if v.Name == "" {
		return errors.New("nome é obrigatório")
	}
	if v.LastName == "" {
		return errors.New("sobrenome é obrigatório")
	}
	if v.Email == "" {
		return errors.New("email é obrigatório")
	}
	if v.Phone == "" {
		return errors.New("telefone é obrigatório")
	}
	return nil
}
