package repository

import (
	"vetblock/internal/db/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (r *UserRepository) CreateUserTable() {
	r.Db.AutoMigrate(&model.User{})
}

func (r *UserRepository) SaveUser(user *model.User) {
	r.Db.Create(user)
}

func (r *UserRepository) FindUserById(id int) *model.User {
	var user model.User
	r.Db.First(&user, id)
	return &user
}

func (r *UserRepository) FindAllUsers() []model.User {
	var users []model.User
	r.Db.Find(&users)
	return users
}

func (r *UserRepository) FindUserByEmail(email string) *model.User {
	var user model.User
	r.Db.Where("email = ?", email).First(&user)
	return &user
}

func (r *UserRepository) FindUserByUsername(username string) *model.User {
	var user model.User
	r.Db.Where("username = ?", username).First(&user)
	return &user
}

