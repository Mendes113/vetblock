package db

import (
	"fmt"
	"vetblock/internal/db/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func NewDb() *gorm.DB {
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "vetblock", "vet113password", "vetblock")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&model.User{}, model.Animal{}, model.Hospitalization{}, model.Consultation{}, model.ConsultationHistory{}, model.Veterinary{})

    return db
}