package model

import "errors"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
}

func NewUser(id int64, username, password, email string) *User {
	return &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
}

func (u *User) GetID() int64 {
	return u.ID
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetID(id int64) {
	u.ID = id
}

func (u *User) SetUsername(username string) {
	u.Username = username
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) Update(username, password, email string) {
	u.Username = username
	u.Password = password
	u.Email = email
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

func (u *User) ValidatePassword(password string) error {
	if u.Password != password {
		return errors.New("invalid password")
	}
	return nil
}

func (u *User) ValidateEmail(email string) error {
	if u.Email != email {
		return errors.New("invalid email")
	}
	return nil
}

func (u *User) ValidateUsername(username string) error {
	if u.Username != username {
		return errors.New("invalid username")
	}
	return nil
}
