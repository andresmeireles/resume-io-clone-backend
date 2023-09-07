package middleware

import (
	"net/http"
	"strings"

	"github.com/andresmeireles/resume/internal/db/sql"
	"github.com/andresmeireles/resume/internal/db/sql/model"
	"github.com/go-chi/render"
)

func isLogged(r *http.Request) bool {
	db := sql.GetDB()
	auth := r.Header.Get("Authorization")
	auth = strings.Replace(auth, "Bearer ", "", 1)

	if auth == "" {
		return false
	}

	return validateAuth(db, auth)
}

func validateAuth(db sql.DBInterface, token string) bool {
	user := &model.User{}

	err := db.GetOneBy(user, "hash", token)

	if err != nil {
		return false
	}

	if user.IsHashExpired() {
		return false
	}
	return true
}

func MustBeLogged(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isLogged := isLogged(r)

		if !isLogged {
			render.JSON(w, r, map[string]interface{}{
				"error": "Unauthorized",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RedirectToHome(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isLogged := isLogged(r)
		currentRoute := r.URL.Path

		if isLogged && currentRoute == "/login" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
