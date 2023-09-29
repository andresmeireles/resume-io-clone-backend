package routes

import (
	"fmt"
	"net/http"

	"github.com/andresmeireles/resume/internal/db/sql"
	"github.com/andresmeireles/resume/internal/db/sql/model"
	"github.com/andresmeireles/resume/web/controller/resume"
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

	r.Get("/resume", resume.ResumeData)

	r.Post("/skill", resume.AddSkill)
	r.Put("/skill", resume.UpdateSkill)

	r.Post("/social", resume.AddSocial)
	r.Put("/social", resume.UpdateSocial)

	r.Post("/education", resume.AddEducation)
	r.Put("/education", resume.UpdateEducation)

	r.Post("/experience", resume.AddExperience)
	r.Put("/experience", resume.UpdateExperience)

	r.Post("/personal", resume.AddPersonalData)
}
