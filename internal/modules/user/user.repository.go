package user

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	DB *gorm.DB
}

func(r* Repository) Create(user *User) error {
	return r.DB.Create(user).Error
}

func(r* Repository) Update(userId int, updates map[string]any) error{
	err := r.DB.
		Model(&User{}).
		Where("id = ?", userId).
		Updates(updates).
		Error

	return err
}

func(r* Repository) Delete(userId int) error {
res := r.DB.Delete(&User{}, userId)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil}

func(r* Repository) Read(userId int) (*User, error) {
	user := &User{}
	err := r.DB.
		Where("id = ?", userId).
		First(user).
		Error

	return user, err
}

func(r* Repository) ReadByMail(email string) (*User, error) {
	user := &User{}
	err := r.DB.
		Where("email = ?", email).
		First(user).
		Error

	return user, err
}


func(r *Repository) Debit(userId int, amount int) error {
	return r.DB.Transaction(func (tx *gorm.DB) error {
		user := &User{}

		err := tx.Model(user).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", userId).
			First(user).
			Error; 
			
		if err != nil {
			return errors.New("Error founding user account.\n"+err.Error())
		}
		
		updated_balance := *user.Balance - amount
		
		err = tx.Model(&User{}).Where("id=?", userId).Update("balance", updated_balance).Error
		
		if err != nil {
			return errors.New("Error making transaction.\n"+err.Error())
		}

		return err
	})
}

func(r *Repository) Credit(userId int, amount int) error {
	return r.DB.Transaction(func (tx *gorm.DB) error {
		user := &User{}

		err := tx.Model(user).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", userId).
			First(user).
			Error; 
		
		if err != nil {
			return errors.New("Error founding user account.\n"+err.Error())
		}
		
		updated_balance := *user.Balance + amount
		
		err = tx.Model(&User{}).Where("id=?", userId).Update("balance", updated_balance).Error
		
		if err != nil {
			return errors.New("Error making transaction.\n"+err.Error())
		}

		return err
	})
}