package handlers

import (
	"vetblock/internal/db/model"
	"vetblock/internal/service"

	"github.com/gofiber/fiber/v2"
)


func AddAnimalTransactionHandler(c *fiber.Ctx) error {
	var animal model.Animal
	if err := c.BodyParser(&animal); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	// Exemplo de dados do remetente e destinatário
	sender := "System" // Alterar conforme necessário
	receiver := "User" // Alterar conforme necessário
	amount := 0.0      // Alterar conforme necessário

	err := service.AddAnimalTransaction(animal, sender, receiver, amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to add animal transaction")
	}

	return c.Status(fiber.StatusCreated).JSON(animal)
}

func GetAnimalByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id") // Obtém o ID do animal da URL

	animal, err := service.GetAnimalByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.JSON(animal)
}

// Atualiza as informações de um animal
func UpdateAnimal(c *fiber.Ctx) error {
	id := c.Params("id")
	var animal model.Animal
	if err := c.BodyParser(&animal); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	if err := service.UpdateAnimal(id, animal); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to update animal")
	}

	return c.Status(fiber.StatusOK).JSON(animal)
}

// Exclui um animal
func DeleteAnimal(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := service.DeleteAnimal(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to delete animal")
	}

	return c.Status(fiber.StatusNoContent).SendString("Animal deleted")
}
