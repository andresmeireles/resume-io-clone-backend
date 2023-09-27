package controller

import (
	"encoding/json"
	"net/http"

	"github.com/andresmeireles/resume/internal/db/sql"
	"github.com/andresmeireles/resume/internal/db/sql/model"
	"github.com/andresmeireles/resume/internal/service/auth"
	"github.com/go-chi/render"
)

type loginRequest struct {
	Email    string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	var body loginRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := auth.LoginWithEmail(sql.GetDB(), body.Email, body.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, token)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	hash := r.Header.Get("Authorization")
	hash = hash[7:]
	user := &model.User{}

	err := sql.GetDB().GetOneBy(user, "hash", hash)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = auth.Logout(sql.GetDB(), user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"message": "Logged out",
	})
}
