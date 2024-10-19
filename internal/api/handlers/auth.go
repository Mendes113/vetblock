package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"vetblock/internal/db/model"
	"vetblock/internal/service"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenRequest struct {
	IDToken string `json:"idToken"`
}

type LoginResponse struct {
	IDToken string `json:"idToken"`
}

var firebaseAuth *auth.Client

func init() {
	// Inicializando o Firebase Admin SDK
	opt := option.WithCredentialsFile("vetsys.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	firebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	log.Println("Firebase Auth initialized")
}

func RegisterUser(c *fiber.Ctx) error {
	var req struct {
		UserType  string `json:"user_type"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		CPF       string `json:"cpf,omitempty"`
		CRMV      string `json:"crmv,omitempty"`
		Phone     string `json:"phone"`
		Name      string `json:"name"`
		Address   string `json:"address,omitempty"`
		Specialty string `json:"specialty,omitempty"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Criar usuário no Firebase
	params := (&auth.UserToCreate{}).
		Email(req.Email).
		Password(req.Password)

	userRecord, err := firebaseAuth.CreateUser(context.Background(), params)
	if err != nil {
		log.Printf("Failed to create user in Firebase: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	var dbErr error

	// Criar e salvar o usuário no banco de dados dependendo do tipo
	if req.UserType == "tutor" {
		tutor, err := model.NewTutor(req.CPF, req.Name, req.Email, req.Phone, req.Address, req.Password)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to create tutor",
			})
		}

		// Salvar tutor no banco de dados
		dbErr = service.CreateUser(tutor)

	} else if req.UserType == "veterinarian" {
		vet, err := model.NewVeterinarian(req.CRMV, req.Name, req.Email, req.Phone, req.Specialty)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to create veterinarian",
			})
		}

		// Salvar veterinário no banco de dados
		dbErr = service.CreateUser(vet)

	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user type",
		})
	}

	if dbErr != nil {
		// Se houver erro ao salvar no banco, remover usuário do Firebase
		firebaseAuth.DeleteUser(context.Background(), userRecord.UID)
		log.Printf("Failed to save user to database: %v", dbErr)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to save user to database",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

// Função para autenticar usuários com token recebido do cliente
func Authenticate(c *fiber.Ctx) error {
	var body TokenRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Verificar o ID Token usando Firebase Auth
	token, err := firebaseAuth.VerifyIDToken(context.Background(), body.IDToken)
	if err != nil {
		log.Printf("Failed to verify token: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	fmt.Printf("Authenticated User ID: %s\n", token.UID)
	return c.JSON(fiber.Map{
		"message": "Authenticated successfully",
		"uid":     token.UID,
	})
}

// Rota para criar um novo usuário no Firebase
func SignUp(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	log.Printf("Creating user with email: %s\n", req.Email)

	// Criar um usuário no Firebase
	params := (&auth.UserToCreate{}).
		Email(req.Email).
		Password(req.Password)

	userRecord, err := firebaseAuth.CreateUser(context.Background(), params)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to sign up",
		})
	}

	bytes, _ := json.Marshal(userRecord)
	fmt.Printf("User created successfully: %s\n", bytes)
	return c.JSON(fiber.Map{
		"uid":   userRecord.UID,
		"email": userRecord.Email,
	})
}

// Middleware para autenticação por token Bearer
func Auth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Verificar o token JWT do Firebase
	decodedToken, err := firebaseAuth.VerifyIDToken(context.Background(), token)
	if err != nil {
		log.Printf("Failed to verify token: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	fmt.Println("User Info:", decodedToken.UID)
	return c.Next()
}
