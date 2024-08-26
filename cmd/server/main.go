package main

import (
	"log"
	"vetblock/internal/api"
	
	

	"github.com/gofiber/fiber/v2"
)



func main() {
	
	

	

	// Inicializar o Fiber e as rotas
	app := fiber.New()

	
	// Configurar as rotas, passando o serviço
	api.SetupRoutes(app)

	// Rota de teste
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("VetBlockchain API")
	})

	log.Println("Servidor iniciado na porta 8080...")
	log.Fatal(app.Listen(":8081"))
}

