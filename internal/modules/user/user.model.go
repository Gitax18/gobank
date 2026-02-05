package user

import "gorm.io/gorm"

type User struct {
	ID            uint    `bson:"primary key;autoIncrement" json:"id"`
	Name          *string `json:"name"`
	Number        *int    `json:"number"`
	AccountNumber *int    `json:"account_number"`
	Balance       *int  `json:"balance"`
}

func MigrateUser(db *gorm.DB) error{
	err := db.AutoMigrate(&User{})
	return err
}