package main

import (
	"log"
	"vetblock/internal/api"
	"vetblock/internal/blockchain"
	"vetblock/internal/db"
	"vetblock/internal/db/repository"
	"vetblock/internal/network"
	"vetblock/internal/service"

	"github.com/gofiber/fiber/v2"
)



func main() {
	
    database := db.NewDb()
	// Inicializar o repositório
	animalRepo := &repository.AnimalRepository{Db: database}

	// Inicializar o serviço
	srv := service.NewService(animalRepo)

	// Inicializar o Fiber e as rotas
	app := fiber.New()

	// Inicializar a blockchain com o bloco gênese
	blockchain.InitializeBlockchain()
	network.StartServer()

	// Configurar as rotas, passando o serviço
	api.SetupRoutes(app, srv)

	// Rota de teste
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("VetBlockchain API")
	})

	log.Println("Servidor iniciado na porta 8080...")
	log.Fatal(app.Listen(":8082"))
}

