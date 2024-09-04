package api

import (
	"vetblock/internal/api/handlers"
	"vetblock/internal/db/repository"
	"vetblock/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/register", handlers.SignUp)
	app.Post("/api/login", handlers.SignIn	)
	//middleware
	
	protected := app.Group("/api/v1")
	protected.Use(handlers.Auth)
	// Rotas para Animais
	protected.Post("/animals", handlers.AddAnimalHandler())
	protected.Get("/animals", handlers.GetAllAnimalsHandler())
	protected.Post("/animals/dosage", handlers.AddDosageHandler(
		service.NewDosageService(
		repository.NewDosageRepository(
			repository.GetDB(),
		))))
	// protected.Get("/animals/:id", handlers.GetAnimalByIDHandler())
	protected.Delete("/animals/:id", handlers.DeleteAnimalHandler())

	// Rotas para Consultas
	protected.Post("/consultations", handlers.AddConsultationHandler())

	// Rotas para Veterinários
	protected.Post("/veterinaries", handlers.AddVeterinaryHandler())

	// Rotas para Medicamentos
	protected.Post("/medications", handlers.AddMedicationHandler())
	protected.Get("/medications/:id", handlers.GetMedicationByIDHandler())
	protected.Delete("/medications/:id", handlers.DeleteMedicationHandler())
	protected.Put("/medications", handlers.UpdateMedicationHandler())
	protected.Get("/medications", handlers.GetAllMedicationsHandler())
	protected.Get("/medications/closest-expiration", handlers.GetMedicationClosestExpirationDateHandler())
	protected.Get("/medications/expired", handlers.GetExpiredMedicationsHandler())
	protected.Get("/medications/will-expire-in/:days", handlers.GetMedicationsWillExpireInDaysHandler())
	protected.Get("/medications/batch-number/:batch_number", handlers.GetMedicationByBatchNumberHandler())
	protected.Get("/medications/name/:name", handlers.GetMedicationByNameHandler())
	protected.Get("/medications/active-substance/:active_substance", handlers.GetMedicationByActiveSubstanceHandler())
}
