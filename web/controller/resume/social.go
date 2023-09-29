package resume

import (
	"net/http"

	userCollection "github.com/andresmeireles/resume/internal/db/nosql/collection/user"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
	"github.com/andresmeireles/resume/internal/utils/request"
	"github.com/go-chi/render"
)

func AddSocial(w http.ResponseWriter, r *http.Request) {
	usr, err := request.GetUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := request.ParseBody[schema.Social](r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userCollection := userCollection.Instance()

	_, err = userCollection.AddSocial(int(*usr.Id), body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"message": "Social added",
	})
}

func UpdateSocial(w http.ResponseWriter, r *http.Request) {
	usr, err := request.GetUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := request.ParseBody[schema.Social](r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userCollection := userCollection.Instance()

	err = userCollection.UpdateSocial(int(*usr.Id), body)

	if err != nil {
		render.JSON(w, r, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"message": "Social updated",
	})
}
