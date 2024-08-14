package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func NewDb() *gorm.DB {
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "vetblock", "vetblock", "vet113password")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
    if err != nil {
        panic("failed to connect database")
    }

    return db
}