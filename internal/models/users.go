package models

import "database/sql"

type userModel struct {
	DB *sql.DB
}

func (u *userModel) Insert(name, email, password string) error {
	return nil
}

func (u *userModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (u *userModel) Exists(email string) (bool, error) {
	return false, nil
}
