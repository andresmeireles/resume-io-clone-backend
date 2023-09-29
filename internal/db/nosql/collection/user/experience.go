package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
)

func (u *User) getExperiencesCollection(userId int) *firestore.CollectionRef {
	return u.getUserCollection(userId).Collection(experience)
}

func (u User) GetExperiences(userId int) ([]schema.Experience, error) {
	experiences := []schema.Experience{}
	firestoreData, err := u.getExperiencesCollection(userId).Documents(context.TODO()).GetAll()
	if err != nil {
		return []schema.Experience{}, nil
	}
	for _, d := range firestoreData {
		experienceData := d.Data()
		endDate, _ := experienceData["end_date"].(string)
		experiences = append(experiences, schema.Experience{
			Job:         experienceData["job"].(string),
			Company:     experienceData["company"].(string),
			StartDate:   experienceData["start_date"].(string),
			EndDate:     &endDate,
			Description: experienceData["description"].(string),
			Hide:        experienceData["hide"].(bool),
		})
	}

	return experiences, nil
}

func (u *User) AddExperience(userId int, experience schema.Experience) (*firestore.WriteResult, error) {
	job := strings.ReplaceAll(experience.Job, " ", "-")
	company := strings.ReplaceAll(experience.Company, " ", "-")

	return u.getExperiencesCollection(userId).Doc(fmt.Sprintf("%s-%s", job, company)).Set(context.TODO(), experience.ToMap())
}

func (u *User) UpdateExperience(userId int, experience schema.Experience) error {
	doc := u.getExperiencesCollection(userId).Doc(fmt.Sprintf("%s-%s", experience.Job, experience.Company))
	getDoc, err := doc.Get(context.TODO())
	if err != nil {
		return errors.New(fmt.Sprintf("Job %s in %s not found", experience.Job, experience.Company))
	}
	if getDoc.Exists() {
		doc.Set(context.TODO(), experience.ToMap())
		return nil
	}

	return errors.New("Experience not found")
}
