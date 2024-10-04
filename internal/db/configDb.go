package db

import (
	"fmt"
	"log"
	"os"
	"vetblock/internal/db/model"

	"github.com/joho/godotenv"
	supabase "github.com/lengzuo/supa"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func GetDB() *gorm.DB {
	return NewDb()
}

func NewDb() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "vetblock", "vet113password", "vetblock")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Verifica o retorno de erro da migração
	errMigrate := db.AutoMigrate(&model.User{}, &model.Animal{}, &model.Hospitalization{}, &model.Consultation{}, &model.ConsultationHistory{}, &model.Veterinary{}, &model.Medication{}, &model.Dosage{})
	if errMigrate != nil {
		log.Fatalf("failed to auto migrate: %v", errMigrate)
	}

	return db
}

func Supa() *supabase.Client{
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
		return nil
	}

	log.Println("Initializing Supabase client")
	log.Println("API Key:", os.Getenv("SUPA_API"))
	log.Println("Project Ref:", os.Getenv("PROJECT_REF"))

	conf := supabase.Config{
		ApiKey:     os.Getenv("SUPA_API"),
		ProjectRef: os.Getenv("PROJECT_REF"),
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
