package api

import (
	"vetblock/internal/api/handlers"
	// "vetblock/internal/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	protected := app.Group("/api/v1")

	
	
	protected.Post("/animals", handlers.AddAnimalHandler())
	// protected.Get("/animals/:id", handlers.GetAnimalByIDHandler())
	protected.Delete("/animals/:id", handlers.DeleteAnimalHandler())
	protected.Post("/consultations", handlers.AddConsultationHandler())
	protected.Post("/veterinaries", handlers.AddVeterinaryHandler())


}