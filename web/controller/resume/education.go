package resume

import (
	"net/http"

	userCollection "github.com/andresmeireles/resume/internal/db/nosql/collection/user"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
	"github.com/andresmeireles/resume/internal/utils/request"
	"github.com/go-chi/render"
)

func AddEducation(w http.ResponseWriter, r *http.Request) {
	user, err := request.GetUserFromRequest(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := request.ParseBody[schema.Education](r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userCollection := userCollection.Instance()

	_, err = userCollection.AddEducation(int(*user.Id), body)

	render.JSON(w, r, map[string]interface{}{
		"message": "Education added",
	})
}

func UpdateEducation(w http.ResponseWriter, r *http.Request) {
	user, err := request.GetUserFromRequest(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := request.ParseBody[schema.Education](r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userCollection := userCollection.Instance()

	err = userCollection.UpdateEducation(int(*user.Id), body)

	render.JSON(w, r, map[string]interface{}{
		"message": "Education updated",
	})
}
