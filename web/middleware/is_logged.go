package middleware

import (
	"net/http"
	"strings"
)

func isLogged(r *http.Request) bool {
	auth := r.Header.Get("Authorization")
	auth = strings.Replace(auth, "Bearer ", "", 1)
	if auth == "" {
		return false
	}
	return validateAuth(auth)
}

func validateAuth(token string) bool {
	return token != ""
}

func MustBeLogged(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		is_logged := isLogged(r)

		if !is_logged {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
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
