package api

import (
	"vetblock/internal/api/handlers"
	"vetblock/internal/db"
	"vetblock/internal/db/repository"
	"vetblock/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	// Middleware CORS para permitir requisições de outros domínios
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Ajuste para origens específicas em produção por segurança
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Rotas de autenticação
	app.Post("/api/register", handlers.SignUp)
	app.Post("/api/login", handlers.Authenticate)

	// Configurar grupo de rotas protegidas
	protected := app.Group("/api/v1")
	protected.Use(handlers.Auth) // Adicionando middleware de autenticação para rotas protegidas

	// Rotas para Animais
	protected.Post("/animals", handlers.AddAnimalHandler())
	protected.Get("/animals", handlers.GetAllAnimalsHandler())
	protected.Get("/animals/:id", handlers.GetAnimalByIDHandler())
	protected.Post("/animals/dosage", handlers.AddDosageHandler(
		service.NewDosageService(
			repository.NewDosageRepository(
				repository.GetDB(),
			))))
	protected.Delete("/animals/:id", handlers.DeleteAnimalHandler())

	// Rotas para Consultas
	protected.Post("/consultations", handlers.AddConsultationHandler(repository.NewConsultationRepository(db.GetDB())))
	protected.Get("/veterinary/:crvm/next-consultation", handlers.GetNextConsultationHandler(repository.NewConsultationRepository(db.GetDB())))
	protected.Get("/consultations/:crvm", handlers.GetAllConsultationsByVeterinaryHandler(repository.NewConsultationRepository(db.GetDB())))
	protected.Post("/veterinaries", handlers.AddVeterinaryHandler())
	protected.Get("/consultations/patient/:animal_id", handlers.GetConsultsByAnimalIDHandler(repository.NewConsultationRepository(db.GetDB())))

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

	// Rotas para Imagens
	imageRepo := repository.NewImageRepository()          // Repositório de imagens
	imageService := service.NewImageService(imageRepo)    // Serviço de imagens
	imageHandler := handlers.NewImageHandler(imageService) // Handler de imagens

	protected.Get("/animals/:id/image", imageHandler.GetImageByIDHandler)
}
