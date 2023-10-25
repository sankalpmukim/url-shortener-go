package database

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
	// SignUpUser(user User) error
	Close() error
}
