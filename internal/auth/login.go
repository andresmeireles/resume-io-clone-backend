package auth

import (
	"errors"

	"github.com/andresmeireles/resume/db/sql"
)

func LoginWithEmail(email, password string) (*sql.User, error) {
	user, err := sql.User{}.FindByEmail(email)
	if err != nil {
		panic(err)
	}
	if user.VerifyPassword(password) {
		return nil, errors.New("Invalid password")
	}
	return user, nil
}
