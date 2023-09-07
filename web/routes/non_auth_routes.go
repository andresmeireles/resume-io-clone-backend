package routes

import (
	"net/http"

	"github.com/andresmeireles/resume/web/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func nonAuthRoutes(r chi.Router) {
	r.Post("/login", controller.Login)
	r.Get("/logout", controller.Logout)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"message": "Hello World",
		}
		render.JSON(w, r, response)
	})
}
