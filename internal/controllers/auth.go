package controllers

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"os"

	"github.com/google/uuid"

	"github.com/sankalpmukim/url-shortener-go/internal/database"
)

func CheckUserCredentials(email string, password string) bool {
	return email == "sankalpmukim@gmail.com" || password == "root"
}

func computeSHA512(input string) string {
	hash := sha512.Sum512([]byte(input))
	return hex.EncodeToString(hash[:])
}

func CreateSecretPayload(email string) string {
	// create secret token
	secret := os.Getenv("SECRET")
	secretPayload := computeSHA512(secret + email)
	combinedPayload := secretPayload + "." + email
	// base64 encode the combined payload
	encodedPayload := base64.URLEncoding.EncodeToString([]byte(combinedPayload))

	return encodedPayload
}

func CheckIfEmailExists(email string) bool {
	emails := []string{"sankalpmukim@gmail.com"}
	for _, e := range emails {
		if e == email {
			return false
		}
	}
	return true
}

func CreateUser(fullName, email, password string) (database.User, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return database.User{}, err
	}
	return database.User{
		ID:       u.String(),
		FullName: fullName,
		Email:    email,
		Password: password,
	}, nil
}
