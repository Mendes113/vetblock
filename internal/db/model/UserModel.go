package model



type User struct {
    ID       uint   `json:"id" gorm:"primary_key"`
    Email    string `json:"email" validate:"required,email"`
    Phone    string `json:"phone" validate:"required,min=10,max=15"`
    Password string `json:"-" validate:"required,min=8"`
}

