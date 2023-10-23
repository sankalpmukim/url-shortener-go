package controllers

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/sankalpmukim/url-shortener-go/internal/database"
	"github.com/sankalpmukim/url-shortener-go/internal/lib"
	"github.com/sankalpmukim/url-shortener-go/pkg/logs"
)

type contextKey string

const emailKey contextKey = "email"

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

func createSecretPayload(email string) string {
	// create secret token
	secret := os.Getenv("SECRET")
	secretPayload := lib.ComputeSHA512(secret + email)

	return secretPayload
}

func CreateEncodedPayload(email string) string {
	// create secret token
	secretPayload := createSecretPayload(email)
	combinedPayload := secretPayload + "." + email
	// base64 encode the combined payload
	encodedPayload := base64.URLEncoding.EncodeToString([]byte(combinedPayload))

	return encodedPayload
}

func GetEmailFromPayload(payload string) (string, error) {
	// base64 decode the payload
	decodedPayload, err := base64.URLEncoding.DecodeString(payload)
	if err != nil {
		logs.Error("Error decoding base64 payload", err)
		return "", errors.New("invalid payload")
	}

	// split the payload
	payloadParts := strings.Split(string(decodedPayload), ".")
	if len(payloadParts) != 3 {
		logs.Error("Wrong number of parts found in auth cookie payload")
		logs.Info("Payload parts", payloadParts)
		return "", errors.New("invalid payload")
	}

	// get the email
	email := payloadParts[1] + "." + payloadParts[2]

	// verify the payload
	secretPayload := createSecretPayload(email)

	if secretPayload != payloadParts[0] {
		logs.Error("Auth cookie payload verification failed")
		logs.Info("Payload parts", payloadParts)
		logs.Info("Email", email)
		logs.Info("Secret payload", secretPayload)
		return "", errors.New("invalid payload")
	}

	// return the email
	return email, nil
}

func SetEmailInContext(r *http.Request, email string) context.Context {
	return context.WithValue(r.Context(), emailKey, email)
}

func GetMailFromContext(r *http.Request) string {
	return r.Context().Value(emailKey).(string)
}
