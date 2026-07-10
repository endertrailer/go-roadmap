package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(byte), err
}

func VerifyPassword(password string, hash string) bool {
	result := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return result == nil
}
