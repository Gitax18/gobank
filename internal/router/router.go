package router

import (
	"github.com/Gitax18/gobank/internal/modules/transaction"
	"github.com/Gitax18/gobank/internal/modules/user"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB){
	user.UserRoutes(app, db)
	transaction.TransactionRoutes(app, db)
}