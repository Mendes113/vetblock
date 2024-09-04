package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"vetblock/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/lengzuo/supa/dto"
)

var supaClient = db.Supa()

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	IDToken string `json:"idToken"`
}

func Auth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	resp, err := supaClient.Auth.User(context.Background(), token)
	if err != nil {
		log.Printf("failed to get user: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	fmt.Println("User Info:", resp.Email)
	return c.Next()
}

func SignUp(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	log.Println("req", req)
	resp, err := supaClient.Auth.SignUp(context.Background(), dto.SignUpRequest{
		Email: req.Email,
		Password: req.Password,
	})
	log.Println("resp", resp)
	if err != nil {
		log.Printf("failed to sign up: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to sign up",
		})
	}

	bytes, _ := json.Marshal(resp)
	fmt.Printf("sign up success: %s", bytes)
	return c.JSON(resp)
}

func SignIn(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Ajuste o método para fazer login conforme a documentação da biblioteca
	resp, err := supaClient.Auth.SignInWithPassword(context.Background(), dto.SignInRequest{	
		Email: req.Email,
		Password: req.Password,
	})
	if err != nil {
		log.Printf("failed to sign in with password: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to sign in",
		})
	}

	bytes, _ := json.Marshal(resp)
	fmt.Printf("sign in with password success: %s\n", bytes)
	return c.JSON(resp)
}
