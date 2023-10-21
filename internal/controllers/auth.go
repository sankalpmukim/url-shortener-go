package controllers

import (
	"encoding/base64"
	"os"

	"github.com/sankalpmukim/url-shortener-go/internal/database"
	"github.com/sankalpmukim/url-shortener-go/internal/lib"
)

func CheckUserCredentials(email string, password string) bool {
	// get user from database
	user, err := database.DB.GetUserByEmail(email)
	if err != nil {
		return false
	}

	// compare password
	if lib.ComparePasswords(user.Password, password) {
		return true
	}

	return false
}

func CreateSecretPayload(email string) string {
	// create secret token
	secret := os.Getenv("SECRET")
	secretPayload := lib.ComputeSHA512(secret + email)
	combinedPayload := secretPayload + "." + email
	// base64 encode the combined payload
	encodedPayload := base64.URLEncoding.EncodeToString([]byte(combinedPayload))

	return encodedPayload
}
