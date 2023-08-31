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

	port := utils.GetEnv("WEB_PORT")
	if port == "" {
		panic("WEB_PORT is not set")
	}

	http.ListenAndServe(":"+port.(string), r)
}
