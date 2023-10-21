package database

type DBInterface interface {
	// GetUserData(userID string) (*User, error)
	// SignUpUser(user User) error
	Close() error
}
