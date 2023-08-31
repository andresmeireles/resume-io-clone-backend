package utils

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, error) {
	if strings.Trim(password, "") == "" {
		return "", errors.New("Password cannot be empty")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
