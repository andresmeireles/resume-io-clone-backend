package routes

import "github.com/go-chi/chi/v5"

func Routes(r *chi.Mux) {
	non_auth_routes(r)
	private_routes(r)
}
