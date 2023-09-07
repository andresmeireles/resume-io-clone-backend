package routes

import (
	"github.com/andresmeireles/resume/web/middleware"
	"github.com/go-chi/chi/v5"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func Routes(r chi.Router) {
	r.Use(chiMiddleware.Logger)
	r.Use(middleware.SetJSON)

	nonAuthRoutes(r)

	r.Group(func(r chi.Router) {
		privateRoutes(r)
	})
}
