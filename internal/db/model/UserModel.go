package model

type UserType string

const (
    TutorType       UserType = "tutor"
    VeterinarianType UserType = "veterinarian"
)

type User struct {
    ID       uint     `json:"id" gorm:"primary_key"`
    Email    string   `json:"email" validate:"required,email"`
    Phone    string   `json:"phone" validate:"required,min=10,max=15"`
    UserType UserType `json:"user_type" gorm:"type:varchar(20)" validate:"required"`
}
