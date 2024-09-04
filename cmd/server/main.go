package main

import (
	"log"
	"vetblock/internal/api"
	

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	
	loadEnv()
	// Inicializar o Fiber e as rotas
	app := fiber.New()

	// Configurar as rotas, passando o serviço
	api.SetupRoutes(app)

	log.Println("Servidor iniciado na porta 8080...")
	log.Fatal(app.Listen(":8081"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println("Environment variables loaded")
}
