package user

import (
	"errors"

	"gorm.io/gorm"
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