package database

import "github.com/google/uuid"

type DBInterface interface {
	// users
	GetUsers() ([]User, error)
	CreateUser(user CreateUser) error
	UserExists(email string) bool
	GetUserByEmail(email string) (User, error)

	// links
	GetLinks() ([]Link, error)
	GetLinkByEndpoint(endpoint string) (Link, error)
	LinkExists(target string) bool
	UpdateLink(oldEndpoint, newEndpoint, target string) error
	CreateLink(link CreateLink) error
	IncrementClicks(endpoint string) error
	GetUserLinks(userID uuid.UUID) ([]Link, error)
	UserLinkExists(userID uuid.UUID, endpoint string) bool
	// SignUpUser(user User) error
	Close() error
}
