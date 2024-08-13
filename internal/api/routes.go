package api

import (
	"vetblock/internal/api/handlers"
	"vetblock/internal/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {


	// app.Post("/register", register)

	// // Rota pública para login
	// app.Post("/login", login)
	protected := app.Group("/api/v1") // Exemplo de prefixo para rotas protegidas
	protected.Use(middleware.AuthMiddleware())


    protected.Get("/blocks", handlers.GetBlockchain)
    protected.Post("/blocks", handlers.AddBlock)
	protected.Get("/blocks/:index/transactions", handlers.GetTransactions)
	protected.Post("/animals", handlers.AddAnimalTransactionHandler)
	protected.Get("/animals/:id", handlers.GetAnimalByIDHandler)


	protected.Post("/consultation/schedule", handlers.ScheduleConsultationHandler)
	protected.Post("/consultation/:id/cancel", handlers.CancelConsultationHandler)
	protected.Post("/consultation/:id/confirm", handlers.ConfirmConsultationHandler)
	protected.Put("/consultation/:id", handlers.UpdateConsultationHandler)
	protected.Get("/consultation/animal/:animal_id", handlers.GetConsultationByAnimalIDHandler)
	protected.Get("/consultation/veterinary/:veterinary_id", handlers.GetConsultationByVeterinaryIDHandler)
}
