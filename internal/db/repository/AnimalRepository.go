package repository

import (
	"log"
	"vetblock/internal/db/model"

	"gorm.io/gorm"
)

type AnimalRepository struct {
	Db *gorm.DB
}


func (r *AnimalRepository) SaveAnimal(animal *model.Animal){
	r.Db.Create(animal)
	log.Print(animal)
	log.Print("Repository Saving Animal")
}