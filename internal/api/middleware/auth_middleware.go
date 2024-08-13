package middleware

import (
	"log"
	"strings"
	"time"
	"vetblock/internal/auth"
	"vetblock/internal/blockchain"

	"github.com/gofiber/fiber/v2"
)


type RateLimiter struct {
	Max        int
	Expiration time.Duration
	Limit 	   func(c *fiber.Ctx) error
}

func New(config RateLimiter) *RateLimiter {
	return &RateLimiter{
		Max:        config.Max,
		Expiration: config.Expiration,
	}
}

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extrair o token do cabeçalho Authorization (ex: "Bearer <token>")
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token não fornecido",
			})
		}

		// Verificar se o cabeçalho está no formato esperado
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Formato de token inválido",
			})
		}
		token := splitToken[1]

		// Validar o token JWT
		_, err := auth.ValidateJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido ou expirado",
			})
		}

		// Token válido, prosseguir com a requisição
		return c.Next()
	}
}
func ValidateRequestBody() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if len(c.Body()) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Corpo da requisição não pode estar vazio",
			})
		}

		// Possíveis validações adicionais

		return c.Next()
	}
}

func LoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Printf("Método: %s, URL: %s", c.Method(), c.OriginalURL())
		return c.Next()
	}
}

func RateLimiterVerify() fiber.Handler {
	// Configurar o Rate Limiter com Fiber (componente adicional ou personalizado)
	limiter := New(RateLimiter{
		Max:        10,          // máximo de requisições
		Expiration: time.Minute, // por minuto
	})
	return limiter.Limit
}

func BlockchainIntegrityMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !blockchain.IsValid() { // Método hipotético para verificar a integridade da blockchain
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Blockchain corrompida",
			})
		}
		return c.Next()
	}
}
