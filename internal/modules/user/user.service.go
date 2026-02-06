package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	r *Repository
}

func (s *Service) CreateUser(email string, password string, name string, number int, account_number int, balance int) error {
	hashBytesPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10); if err != nil{
		return err
	}

	hashPassword := string(hashBytesPassword)
	

	user := User{
		Email: 		   		&email,
		HashedPassword: 	&hashPassword,
		Name:          		&name,
		Number:        		&number,
		AccountNumber: 		&account_number,
		Balance:       		&balance,
	}

	return s.r.Create(&user)
}

func (s *Service) UpdateUser(userid int, name *string, number *int) error {
	updates := make(map[string]any)

	if name != nil {
		updates["name"] = *name
	}

	if number != nil {
		updates["number"] = *number
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	return s.r.Update(userid, updates)
}

func (s *Service) DeleteUser(userid int) error {
	return s.r.Delete(userid)
}

func (s *Service) ReadUser(userid int) (*User, error) {
	return s.r.Read(userid)
}

func (s *Service) ReadUserByMail(email string) (*User, error) {
	return s.r.ReadByMail(email)
}
