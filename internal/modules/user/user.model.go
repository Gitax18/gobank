package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	CreatedAt	  		time.Time		`json:"-"`
	ID            		uint    		`gorm:"primarykey;autoIncrement" json:"id"`
	Email 	  			*string 		`gorm:"unique" json:"email"`
	HashedPassword	 	*string 		`gorm:"unique" json:"password"`
	Name          		*string 		`json:"name"`
	Number        		*int    		`gorm:"unique" json:"number"`
	AccountNumber 		*int    		`gorm:"unique" json:"account_number"`
	Balance       		*int	  		`json:"balance"`
}

func MigrateUser(db *gorm.DB) error{
	err := db.AutoMigrate(&User{})
	return err
}