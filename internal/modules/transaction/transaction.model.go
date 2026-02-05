package transaction

import "gorm.io/gorm"

type Transaction struct {
	ID           uint    	`bson:"primary key;autoIncrement" json:"id"`
	SenderId     *string 	`json:"sender_id"`
	ReceiverId   *int    	`json:"receiver_id"`
	Amount       *int 	`json:"amount"`
}

func MigrateTransaction(db *gorm.DB) error{
	err := db.AutoMigrate(&Transaction{})
	return err
}