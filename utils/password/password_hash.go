package password

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, error) {
	if strings.Trim(password, "") == "" {
		return "", errors.New("Password must not be empty")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password[:72]), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
