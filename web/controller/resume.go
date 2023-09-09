package controller

import (
	"net/http"

	"github.com/andresmeireles/resume/internal/db/nosql"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
	sk "github.com/andresmeireles/resume/internal/service/resume"
	"github.com/andresmeireles/resume/utils/request"
	"github.com/go-chi/render"
)

func AddSkill(w http.ResponseWriter, r *http.Request) {
	usr, err := request.GetUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := request.ParseBody[schema.Skill](r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := nosql.GetNoSqlClient()

	_, err = sk.Add(&client, *usr, body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"message": "Skill added",
	})
}
