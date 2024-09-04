package db

import (
	"fmt"
	"vetblock/internal/db/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
    supabase "github.com/lengzuo/supa"
)

func NewDb() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "vetblock", "vet113password", "vetblock")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{}, model.Animal{}, model.Hospitalization{}, model.Consultation{}, model.ConsultationHistory{}, model.Veterinary{}, model.Medication{}, model.Dosage{})

	return db
}

func Supa() *supabase.Client{
	conf := supabase.Config{
		ApiKey:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Imdjd3hrcHZ5aGtmaHVka2NueHlhIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MjUzODA0NzQsImV4cCI6MjA0MDk1NjQ3NH0.AiD4dUBpAK34c83A0kstm2V4LKX7zajHTVqYo9cSHzg",
		ProjectRef: "gcwxkpvyhkfhudkcnxya",
		Debug:      true,
	}

	supaClient, err := supabase.New(conf)
	if err != nil {
		fmt.Println("failed to initialize Supabase client:", err)
		return nil
	}
	fmt.Println("Supabase client initialized")
    
    return supaClient
}
