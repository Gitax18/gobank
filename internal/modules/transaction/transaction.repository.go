package transaction

import (
	"errors"

	"github.com/Gitax18/gobank/internal/modules/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Create(transaction *Transaction) error {
	// initiate the tsx
	return r.DB.Transaction(func(tx *gorm.DB)error {
		// retrivie sender
		var senderUser user.User
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&senderUser, transaction.SenderId).Error; if err != nil {
			return err
		}
		// validat sender balance for amount deduction
		if *senderUser.Balance <= transaction.Amount {
			return errors.New("Insufficient balance")
		}
		// decrease the balance from amount
		updatedBalance := *senderUser.Balance - transaction.Amount;

		// update the sender details
		err = tx.Model(&senderUser).Update("balance", updatedBalance).Error; if err != nil {
			return err
		}
		// retrieve receiver
		var receiverUser user.User
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&receiverUser, transaction.ReceiverId).Error; if err != nil {
			return err
		}
		// increment the amount
		updatedBalance = *receiverUser.Balance + transaction.Amount;

		// update the receiver details
		err = tx.Model(&receiverUser).Update("balance", updatedBalance).Error; if err != nil {
			return err
		}

		err = tx.Create(transaction).Error; if (err != nil ){
			return err
		}

		// commit the transaction
		return nil

	})
}
