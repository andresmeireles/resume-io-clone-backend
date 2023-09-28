package user

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
)

func (u User) getEducationCollection(userId int) *firestore.CollectionRef {
	return u.getUserCollection(userId).Collection("education")
}

func (u User) AddEducation(userId int, education schema.Education) (*firestore.WriteResult, error) {
	return u.getEducationCollection(userId).Doc("education").Set(context.TODO(), education.ToMap())
}

func (u User) GetEducation(userId int) ([]schema.Education, error) {
	data, err := u.getEducationCollection(userId).Documents(context.TODO()).GetAll()
	if err != nil {
		return nil, err
	}

	educations := []schema.Education{}

	for _, data := range data {
		educationData := data.Data()
		educations = append(educations, schema.Education{
			Degree:      educationData["degree"].(string),
			Institution: educationData["institution"].(string),
			StartDate:   educationData["startDate"].(string),
			EndDate:     educationData["endDate"].(string),
			Description: educationData["description"].(string),
			Hide:        educationData["hide"].(bool),
		})
	}

	return educations, err
}

func (u User) UpdateEducation(userId int, education schema.Education) error {
	doc := u.getEducationCollection(userId).Doc("education")
	getDoc, err := doc.Get(context.TODO())
	if err != nil {
		return err
	}
	if getDoc.Exists() {
		doc.Set(context.TODO(), education.ToMap())
		return nil
	}

	return errors.New("Education not found")
}
