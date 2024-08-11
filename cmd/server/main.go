package main

import (
	"log"
	"vetblock/internal/api"
	"vetblock/internal/blockchain"
	"vetblock/internal/network"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // Inicializar a blockchain com o bloco gênese
    blockchain.InitializeBlockchain()
	network.StartServer()
    // Configurar as rotas
    api.SetupRoutes(app)

    // Rota de teste
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("VetBlockchain API")
    })

    log.Println("Servidor iniciado na porta 8080...")
    log.Fatal(app.Listen(":8080"))
}
