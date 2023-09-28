package resume

import (
	"net/http"

	"github.com/andresmeireles/resume/internal/db/nosql/collection"
	"github.com/andresmeireles/resume/internal/utils/request"
	"github.com/go-chi/render"

	userCollection "github.com/andresmeireles/resume/internal/db/nosql/collection/user"
)

func ResumeData(w http.ResponseWriter, r *http.Request) {
	user, err := request.GetUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uCollection := userCollection.Instance()

	resume, err := collection.GetResumeByUserId(*uCollection, int(*user.Id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"data": resume,
	})
}
