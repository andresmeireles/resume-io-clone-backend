package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andresmeireles/resume/utils"
	"github.com/andresmeireles/resume/web/routes"
	"github.com/go-chi/chi/v5"
)

func Run() {
	context.Background()

	r := chi.NewRouter()

	routes.Routes(r)

	port := utils.GetEnv("WEB_PORT")
	if port == "" {
		panic("WEB_PORT is not set!")
	}

	err := http.ListenAndServe(":"+port.(string), r)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Server is running on port:", port)
}
