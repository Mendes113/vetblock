package model

import "strconv"

type Tutor struct {
    CPF      string `json:"cpf" gorm:"type:char(11);primary_key"` // Ajustado para 11 caracteres
    Name     string `json:"name"`
    Email    string `json:"email"`
    Phone    string `json:"phone"`
    Address  string `json:"address"`
    Password string `json:"-"` 
}


func NewTutor(cpf, name, email, phone, address, password string) *Tutor {
	return &Tutor{
		CPF: cpf,
		Name: name,
		Email: email,
		Phone: phone,
		Address: address,
		Password: password,
	}
}

func (t *Tutor) GetCPF() string {
	return t.CPF
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

func (t *Tutor) GetPassword() string {
	return t.Password
}

func (t *Tutor) SetCPF(cpf string) {
	t.CPF = cpf
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

func (t *Tutor) SetPassword(password string) {
	t.Password = password
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
    if secondDigit != int(cpf[10]-'0') {
        return false
    }

    return true
}
