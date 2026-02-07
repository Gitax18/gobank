package transaction

import (
	"github.com/Gitax18/gobank/internal/middleware"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)


func TransactionRoutes(app *fiber.App, db *gorm.DB){
	// dependency injection
	repo := &Repository{DB: db}
	service := &Service{r: repo}
	handler := &Handler{s: service} 

	router := app.Group("/transaction")
	
	router.Post("/", middleware.CheckAuth, handler.POSTCreateTransaction)

}