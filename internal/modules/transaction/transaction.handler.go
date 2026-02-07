package transaction

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	s *Service
}

type TransactionDTO struct {
	SenderId 	int 	
	ReceiverId 	int 	`json:"receiver_id"`
	Amount 		int 	`json:"amount"`
}

func (h *Handler) POSTCreateTransaction(context fiber.Ctx) error {
	var transaction TransactionDTO
	
	err := context.Bind().JSON(&transaction); if err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Bad request", "err": err.Error()})
	}

	// getting id from request local storage (user data was injected through auth middleware)
	t := context.Locals("user").(map[string]any)
	id:= int(t["id"].(float64))
	transaction.SenderId = id
	
	if transaction.ReceiverId <= 0 || transaction.SenderId <= 0 || transaction.Amount <= 0 {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Bad request", "err": "Incorrect data for making transaction"})
	}
	
	err = h.s.CreateTransaction(transaction.SenderId, transaction.ReceiverId, transaction.Amount); if err != nil {
		return context.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Internal server error", "err": err.Error()})
	}
	
	return context.Status(http.StatusOK).JSON(fiber.Map{"message": "Transaction completed successfully"})
}