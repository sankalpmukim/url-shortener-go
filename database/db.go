package database

import (
	"errors"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/sankalpmukim/url-shortener-go/logs"

	_ "github.com/lib/pq" // Driver for Postgres
)

type Database struct {
	conn *sqlx.DB
}

func NewPostgres(connectionString string) (*Database, error) {
	conn, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &Database{conn: conn}, nil
}

func (db *Database) Close() error {
	return db.conn.Close()
}

func Initialize() error {
	connString := os.Getenv("DATABASE_URL")

	if connString == "" {
		return errors.New("DATABASE_URL not set")
	}

	var err error

	DB, err = NewPostgres(connString)
	if err != nil {
		return err
	}
	logs.Info("Connected to database")
	return nil
}

// Ensure our DB struct implements DBInterface
var DB DBInterface = &Database{}
