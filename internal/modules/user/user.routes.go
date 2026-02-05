package user

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)


func UserRoutes(app *fiber.App, db *gorm.DB){
	// dependency injection
	repo := &Repository{DB: db}
	service := &Service{r: repo}
	handler := &Handler{s: service} 

	router := app.Group("/user")
	router.Get("/:id", handler.GETUser)
	router.Post("/create", handler.POSTCreateUser)
	router.Put("/update/:id", handler.PUTUpdateUser)
	router.Delete("/delete/:id", handler.DELETEUser)
}