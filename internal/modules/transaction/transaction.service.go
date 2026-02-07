package transaction

import "errors"

type Service struct {
	r *Repository
}

func (s *Service) CreateTransaction(senderId int, receiverId int, amount int) error {

	if senderId == receiverId {
		return errors.New("You cannot transfer money to yourself")
	}

	transaction := Transaction{
		SenderId:   senderId,
		ReceiverId: receiverId,
		Amount:     amount,
	}

	return s.r.Create(&transaction)
}