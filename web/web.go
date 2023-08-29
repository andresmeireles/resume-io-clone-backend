package web

import (
	"net/http"

	"github.com/andresmeireles/resume/utils"
	"github.com/andresmeireles/resume/web/routes"
	"github.com/go-chi/chi/v5"
)

func Run() {
	r := chi.NewRouter()

	routes.Routes(r)

	http.ListenAndServe(":"+utils.GetEnv("WEB_PORT"), r)
}
