package user

import (
	"context"
	"errors"
	"fmt"

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
		experiences = append(experiences, schema.Experience{
			Job:     experienceData["job"].(string),
			Company: experienceData["company"].(string),
		})
	}

	return experiences, nil
}

func (u *User) AddExperience(userId int, experience schema.Experience) (*firestore.WriteResult, error) {
	return u.getSkills(userId).Doc(fmt.Sprintf("%s-%s", experience.Job, experience.Company)).Set(context.TODO(), experience.ToMap())
}

func (u *User) UpdateExperience(userId int, experience schema.Experience) error {
	doc := u.getSkills(userId).Doc(fmt.Sprintf("%s-%s", experience.Job, experience.Company))
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
