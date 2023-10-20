package database

type DBInterface interface {
	// GetUserData(userID string) (*User, error)
	// SignUpUser(user User) error
	Close() error
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
