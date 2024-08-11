package handlers

import (
	"strconv"
	"vetblock/internal/blockchain"

	"github.com/gofiber/fiber/v2"
)

// Função para obter a blockchain
func GetBlockchain(c *fiber.Ctx) error {
    blockchainJSON, err := blockchain.GetBlockchainAsJSON()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Error retrieving blockchain")
    }
    return c.SendString(blockchainJSON)
}


func AddBlock(c *fiber.Ctx) error {
    difficulty, err := strconv.Atoi(c.Query("difficulty", "2"))
    if err != nil {
        difficulty = 2 // valor padrão
    }

    var transactions []blockchain.Transaction
    if err := c.BodyParser(&transactions); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
    }

    newBlock := blockchain.CreateNewBlock(transactions)

    err = blockchain.AddBlockWithConsensus(newBlock, difficulty)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    return c.Status(fiber.StatusCreated).JSON(newBlock)
}