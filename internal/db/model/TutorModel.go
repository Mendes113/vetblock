package model

import (
	"errors"
	"strconv"
)



type Tutor struct {
    User
    CPFTutor string `json:"cpf_tutor" gorm:"type:char(11);primary_key" validate:"required,len=11,cpf"`
    Name     string `json:"name" validate:"required,min=2,max=100"`
    Address  string `json:"address" validate:"required,min=5,max=255"`
}

type Veterinarian struct {
    User
    CRMV     string `json:"crmv" gorm:"type:char(10);primary_key" validate:"required,len=10"`
    Name     string `json:"name" validate:"required,min=2,max=100"`
    Specialty string `json:"specialty" validate:"required,min=2,max=100"`
}

func NewTutor(cpf, name, email, phone, address, password string) (*Tutor, error) {
    if !isValidCPF(cpf) {
        return nil, errors.New("CPF inválido")
    }
    return &Tutor{
        User: User{
            Email:    email,
            Phone:    phone,
            UserType: TutorType,
        },
        CPFTutor: cpf,
        Name:     name,
        Address:  address,
    }, nil
}

func (t *Tutor) GetCPF() string {
	return t.CPFTutor
}

func (t *Tutor) GetName() string {
	return t.Name
}

func (t *Tutor) GetEmail() string {
	return t.Email
}

func (t *Tutor) GetPhone() string {
	return t.Phone
}

func (t *Tutor) GetAddress() string {
	return t.Address
}


func (t *Tutor) SetCPF(cpf string) {
	t.CPFTutor = cpf
}

func (t *Tutor) SetName(name string) {
	t.Name = name
}

func (t *Tutor) SetEmail(email string) {
	t.Email = email
}

func (t *Tutor) SetPhone(phone string) {
	t.Phone = phone
}

func (t *Tutor) SetAddress(address string) {
	t.Address = address
}

func cleanCPF(cpf string) string {
    cleaned := ""
    for _, char := range cpf {
        if char >= '0' && char <= '9' {
            cleaned += string(char)
        }
    }
    return cleaned
}

// Calcula o dígito verificador
func calculateDigit(cpf string, multiplier int) int {
    sum := 0
    for i, char := range cpf {
        digit, _ := strconv.Atoi(string(char))
        sum += digit * (multiplier - i)
    }
    remainder := sum % 11
    if remainder < 2 {
        return 0
    }
    return 11 - remainder
}

// Valida o CPF
func isValidCPF(cpf string) bool {
    cpf = cleanCPF(cpf)
    
    if len(cpf) != 11 {
        return false
    }

    // Verifica se todos os dígitos são iguais (ex: 111.111.111-11)
    allEqual := true
    for i := 1; i < len(cpf); i++ {
        if cpf[i] != cpf[0] {
            allEqual = false
            break
        }
    }
    if allEqual {
        return false
    }

    // Calcula o primeiro dígito verificador
    firstDigit := calculateDigit(cpf[:9], 10)
    if firstDigit != int(cpf[9]-'0') {
        return false
    }

    // Calcula o segundo dígito verificador
    secondDigit := calculateDigit(cpf[:10], 11)
    return secondDigit == int(cpf[10]-'0')
}
