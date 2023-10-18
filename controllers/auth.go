package controllers

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"os"
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
