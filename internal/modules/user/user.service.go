package user

import "errors"

type Service struct {
	r *Repository
}

func (s *Service) CreateUser(name string, number int, account_number int, balance int) error {
	user := User{
		Name:          &name,
		Number:        &number,
		AccountNumber: &account_number,
		Balance:       &balance,
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
