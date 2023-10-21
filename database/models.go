package database

type User struct {
	ID        string  `json:"id"`
	FullName  string  `json:"full_name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	CreatedAt float64 `json:"created_at"`
	UpdatedAt float64 `json:"updated_at"`
}
