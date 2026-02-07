package main

import (
	"log"
	"os"

	"github.com/Gitax18/gobank/internal/database"
	"github.com/Gitax18/gobank/internal/modules/transaction"
	"github.com/Gitax18/gobank/internal/modules/user"
	"github.com/Gitax18/gobank/internal/router"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	// loading enviroment vars
	err := godotenv.Load(".env"); if err != nil{
		log.Fatal(err)
	}

	// initiating db
	config := &database.Config{
		Host: os.Getenv("DB_HOST"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PWD"),
		Port: os.Getenv("DB_PORT"),
		DBname: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLM"),
	}

	db, err := database.DBConnection(config); if err != nil {
		log.Fatal(err)
	}

	// migrating models to database 
	err = user.MigrateUser(db)
	err = transaction.MigrateTransaction(db)

	// initiang fiber server
	app := fiber.New()

	router.Setup(app, db)

	err = app.Listen(":8080");
	
	if err != nil {
		log.Fatal(err)
	}
}