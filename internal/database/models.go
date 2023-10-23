package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUser struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Link struct {
	Endpoint  string    `json:"endpoint"`
	Target    string    `json:"target"`
	Clicks    int       `json:"clicks"`
	CreatedBy uuid.UUID `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateLink struct {
	Endpoint string `json:"endpoint"`
	Target   string `json:"target"`
}
