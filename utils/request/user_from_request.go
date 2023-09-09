package request

import (
	"net/http"

	"github.com/andresmeireles/resume/internal/db/sql"
	"github.com/andresmeireles/resume/internal/db/sql/model"
)

func GetUserFromRequest(r *http.Request) (*model.User, error) {
	hash := r.Header.Get("Authorization")
	hash = hash[7:]
	user := &model.User{}
	err := sql.GetDB().GetOneBy(user, "hash", hash)

	if err != nil {
		return nil, err
	}

	return user, nil
}
