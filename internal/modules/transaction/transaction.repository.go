package transaction

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(senderId int, receiverId int, amount int){
	
}