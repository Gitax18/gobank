package transaction

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID           	uint   		`gorm:"primarykey;autoIncrement" json:"id"`
	SenderId     	int 		`json:"sender_id"`
	ReceiverId   	int    		`json:"receiver_id"`
	Amount       	int 		`json:"amount"`
}

func MigrateTransaction(db *gorm.DB) error{
	err := db.AutoMigrate(&Transaction{})
	return err
}