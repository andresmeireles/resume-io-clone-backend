package routes

import (
	"github.com/andresmeireles/resume/web/controller"
	"github.com/go-chi/chi/v5"
)

func non_auth_routes(r *chi.Mux) {
	r.Get("/login", controller.Login)
	r.Get("/logout", controller.Logout)
}
