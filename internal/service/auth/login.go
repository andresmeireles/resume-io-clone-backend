package auth

import (
	"errors"
	"time"

	"github.com/andresmeireles/resume/internal/db/sql"
	"github.com/andresmeireles/resume/internal/db/sql/model"
	"golang.org/x/crypto/bcrypt"
)

func LoginWithEmail(db sql.DBInterface, email, password string) (string, error) {
	user := &model.User{}

	defer db.Close()

	e := db.GetOneBy(user, "email", email)

	if e != nil {
		return "", errors.New("Invalid email")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New("Invalid password")
	}

	updatedUser, err := user.SetHash()

	updatedErr := db.UpdateModel(updatedUser)

	if updatedErr != nil {
		return "", updatedErr
	}

	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return *updatedUser.Hash, nil
}

func UserByHash(db sql.DBInterface, hash string) (*model.User, error) {
	user := &model.User{}
	err := db.GetOneBy(user, "hash", hash)
	if err != nil {
		return nil, errors.New("Invalid hash")
	}

	if user.ExpiresAt == nil {
		return nil, errors.New("Invalid hash")
	}
	if time.Now().Unix() > user.ExpiresAt.Unix() {
		return nil, errors.New("Invalid hash")
	}

	return user, nil
}
