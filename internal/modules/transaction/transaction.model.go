package transaction

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	CreatedAt					time.Time		
	ID           				uint   			`gorm:"primarykey;autoIncrement" json:"id"`
	SenderAccountNumber     	*int 			`json:"sender_account_number"`
	ReceiverAccountNumber   	*int    		`json:"receiver_account_number"`
	Amount       				*int 			`json:"amount"`
}

func MigrateTransaction(db *gorm.DB) error{
	err := db.AutoMigrate(&Transaction{})
	return err
}