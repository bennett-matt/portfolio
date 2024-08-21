package models

import "database/sql"

type Models struct {
	Users interface {
		Insert(name, email, password string) error
		Authenticate(email, password string) (int, error)
		Exists(email string) (bool, error)
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users: &userModel{DB: db},
	}
}
