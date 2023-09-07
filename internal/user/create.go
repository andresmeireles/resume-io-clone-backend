package user

import (
	"github.com/andresmeireles/resume/internal/db/sql"
	"github.com/andresmeireles/resume/internal/db/sql/model"
	"github.com/andresmeireles/resume/utils"
)

func Create(db sql.DBInterface, name, password, email string) error {
	hashPassword, err := utils.PasswordHash(password)
	if err != nil {
		return err
	}
	user := model.User{
		Name:     name,
		Email:    email,
		Password: hashPassword,
	}

	err = db.Insert(user)

	return err
}
