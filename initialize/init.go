package initialize

import (
	"github.com/joho/godotenv"
	"github.com/sankalpmukim/url-shortener-go/database"
	"github.com/sankalpmukim/url-shortener-go/logs"
)

func InitAll() error {
	if err := initializeEnv(); err != nil {
		return err
	}
	if err := logs.Initialize(); err != nil {
		return err
	}
	if err := database.Initialize(); err != nil {
		return err
	}
	return nil
}

func initializeEnv() error {
	err := godotenv.Load()
	return err
}
