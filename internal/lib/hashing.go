package lib

import (
	"crypto/sha512"
	"encoding/hex"
)

func ComputeSHA512(input string) string {
	hash := sha512.Sum512([]byte(input))
	return hex.EncodeToString(hash[:])
}

// ComparePasswords compares a hashed password with a plaintext password
func ComparePasswords(hashedPassword string, password string) bool {
	return hashedPassword == ComputeSHA512(password)
}
