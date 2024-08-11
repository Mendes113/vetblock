package api

import (
    "vetblock/internal/api/handlers"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/blocks", handlers.GetBlockchain)
    app.Post("/blocks", handlers.AddBlock)
	app.Get("/blocks/:index/transactions", handlers.GetTransactions)
	app.Post("/animals", handlers.AddAnimalTransactionHandler)
	app.Get("/animals/:id", handlers.GetAnimalByIDHandler)


	app.Post("/consultation/schedule", handlers.ScheduleConsultationHandler)
	app.Post("/consultation/:id/cancel", handlers.CancelConsultationHandler)
	app.Post("/consultation/:id/confirm", handlers.ConfirmConsultationHandler)
	app.Put("/consultation/:id", handlers.UpdateConsultationHandler)
	app.Get("/consultation/animal/:animal_id", handlers.GetConsultationByAnimalIDHandler)
	app.Get("/consultation/veterinary/:veterinary_id", handlers.GetConsultationByVeterinaryIDHandler)
}
