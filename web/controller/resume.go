package controller

import (
	"net/http"

	collection "github.com/andresmeireles/resume/internal/db/nosql/collection/user"
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

	userCollection := collection.Instance()

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

	userCollection := collection.Instance()

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

func AddPersonalData(w http.ResponseWriter, r *http.Request) {
	usr, err := request.GetUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body, err := request.ParseBody[schema.Personal](r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userCollection := collection.Instance()

	_, err = resume.AddPersonal(userCollection, *usr, body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"message": "Personal data updated",
	})
}

func AddExperience(w http.ResponseWriter, r *http.Request) {
	usr, err := request.GetUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body, err := request.ParseBody[schema.Experience](r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userCollection := collection.Instance()

	_, err = userCollection.AddExperience(int(*usr.Id), body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"message": "Experience added",
	})
}

func UpdateExperience(w http.ResponseWriter, r *http.Request) {
	usr, err := request.GetUserFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body, err := request.ParseBody[schema.Personal](r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userCollection := collection.Instance()

	_, err = resume.AddPersonal(userCollection, *usr, body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"message": "Personal data updated",
	})
}
