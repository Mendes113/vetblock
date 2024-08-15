package handlers

import (
	"vetblock/internal/blockchain"

	"github.com/gofiber/fiber/v2"
)

// Função para obter transações de um bloco específico
func GetTransactions(c *fiber.Ctx) error {
    blockIndex := c.Params("index")

    block, err := blockchain.GetBlockByIndex(blockIndex)
    if err != nil {
        return c.Status(fiber.StatusNotFound).SendString("Block not found")
    }

   
    return c.JSON(block.Transactions)
}
