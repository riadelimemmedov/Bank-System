package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// !HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashadPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %v", err)
	}
	return string(hashadPassword), nil
}

// !Checkpassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
