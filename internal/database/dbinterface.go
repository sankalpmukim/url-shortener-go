package database

type DBInterface interface {
	GetUsers() ([]User, error)
	CreateUser(user CreateUser) error
	UserExists(email string) bool
	GetUserByEmail(email string) (User, error)
	// SignUpUser(user User) error
	Close() error
}
