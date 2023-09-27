package auth

import (
	"github.com/andresmeireles/resume/internal/db/sql"
	"github.com/andresmeireles/resume/internal/db/sql/model"
)

func Logout(db sql.DBInterface, user *model.User) error {
	err := db.Update(model.User{}, map[string]interface{}{
		"hash":       nil,
		"expires_at": nil,
	})

	return err
}
