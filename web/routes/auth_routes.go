package routes

import (
	"fmt"
	"net/http"

	"github.com/andresmeireles/resume/internal/db/sql"
	"github.com/andresmeireles/resume/internal/db/sql/model"
	"github.com/andresmeireles/resume/web/controller"
	"github.com/andresmeireles/resume/web/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func privateRoutes(r chi.Router) {
	r.Use(middleware.MustBeLogged)

	r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		hash := r.Header.Get("Authorization")
		hash = hash[7:]
		fmt.Println("hash", hash)
		user := &model.User{}
		err := sql.GetDB().GetOneBy(user, "hash", hash)
		if err != nil {
			render.JSON(w, r, err)
		}
		render.JSON(w, r, user)
	})

	r.Post("/skill", controller.AddSkill)
	r.Put("/skill", controller.UpdateSkill)

	r.Post("/personal", controller.AddPersonalData)
}
