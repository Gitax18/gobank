package user

import (
	"github.com/Gitax18/gobank/internal/middleware"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)


func UserRoutes(app *fiber.App, db *gorm.DB){
	// dependency injection
	repo := &Repository{DB: db}
	service := &Service{r: repo}
	handler := &Handler{s: service} 

	router := app.Group("/user")
	
	router.Post("/login", handler.POSTLoginUser)
	router.Post("/register", handler.POSTCreateUser)
	router.Post("/logout", middleware.CheckAuth, handler.POSTLogoutUser)

	router.Get("", middleware.CheckAuth, handler.GETUser)
	router.Put("/:id", middleware.CheckAuth, handler.PUTUpdateUser)
	router.Delete("/:id", middleware.CheckAuth, handler.DELETEUser)
}