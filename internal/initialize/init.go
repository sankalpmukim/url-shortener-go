package initialize

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sankalpmukim/url-shortener-go/internal/database"
	"github.com/sankalpmukim/url-shortener-go/pkg/logs"
)

func InitAll() error {
	if err := initializeEnv(); err != nil {
		fmt.Println("Error initializing env")
		return err
	}
	if err := logs.Initialize(); err != nil {
		fmt.Println("Error initializing logs")
		return err
	}
	if err := database.Initialize(); err != nil {
		fmt.Println("Error initializing database")
		return err
	}
	return nil
}

func initializeEnv() error {
	err := godotenv.Load()
	return err
}
