package repository

import (
	"vetblock/internal/db"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (r *UserRepository) CreateUserTable() {
	r.Db.AutoMigrate(&db.User{})
}

func (r *UserRepository) SaveUser(user *db.User) {
	r.Db.Create(user)
}

func (r *UserRepository) FindUserById(id int) *db.User {
	var user db.User
	r.Db.First(&user, id)
	return &user
}

func (r *UserRepository) FindAllUsers() []db.User {
	var users []db.User
	r.Db.Find(&users)
	return users
}

func (r *UserRepository) FindUserByEmail(email string) *db.User {
	var user db.User
	r.Db.Where("email = ?", email).First(&user)
	return &user
}

func (r *UserRepository) FindUserByUsername(username string) *db.User {
	var user db.User
	r.Db.Where("username = ?", username).First(&user)
	return &user
}

