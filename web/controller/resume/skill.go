package resume

import (
	"net/http"

	userCollection "github.com/andresmeireles/resume/internal/db/nosql/collection/user"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
	"github.com/andresmeireles/resume/internal/service/resume"
	"github.com/andresmeireles/resume/internal/utils/request"
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

	userCollection := userCollection.Instance()

	_, err = resume.AddSkill(userCollection, *usr, body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"message": "Skill " + body.Name + " added",
	})
}

func UpdateSkill(w http.ResponseWriter, r *http.Request) {
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

	userCollection := userCollection.Instance()

	err = resume.UpdateSkill(userCollection, *usr, body)

	if err != nil {
		render.JSON(w, r, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"message": "Skill " + body.Name + " added",
	})
}
