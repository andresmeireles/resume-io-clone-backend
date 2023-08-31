package routes

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {
	nonAuthRoutes(r)
	privateRoutes(r)
}
