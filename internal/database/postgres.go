package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
	SSLMode  string
}

func DBConnection(config *Config) (*gorm.DB, error) {
	dbcs := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBname, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dbcs), &gorm.Config{}); if err != nil{
		log.Fatal(err);
		return db, err
	}

	return db, err
}