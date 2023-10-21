package database

func (db *Database) GetUsers() ([]User, error) {
	var users []User
	err := db.conn.Select(&users, "SELECT * FROM users")
	return users, err
}

func (db *Database) CreateUser(user CreateUser) error {
	_, err := db.conn.Exec("INSERT INTO users (fullname, email, password) VALUES ($1, $2, $3)",
		user.FullName, user.Email, user.Password)
	return err
}

func (db *Database) UserExists(email string) bool {
	var user User
	err := db.conn.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	return err == nil
}

func (db *Database) GetUserByEmail(email string) (User, error) {
	var user User
	err := db.conn.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	return user, err
}
