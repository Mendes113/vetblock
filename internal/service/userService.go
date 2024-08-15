package service

import (
	"vetblock/internal/db"
	"vetblock/internal/db/model"
)


func CreateUser(model.User) error {
	
	user := model.User{} 
	db.NewDb().Save(&user) // Pass the instance to the Save method
	return nil
}