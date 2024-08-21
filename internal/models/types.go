package models

import "time"

type User struct {
	ID             int64
	Name           string
	Email          string
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ViewData struct {
	Title           string
	Flash           string
	Form            any
	IsAuthenticated bool
	CSRFToken       string
}
