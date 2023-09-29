package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
)

func (u User) getEducationCollection(userId int) *firestore.CollectionRef {
	return u.getUserCollection(userId).Collection("education")
}

func (u User) AddEducation(userId int, education schema.Education) (*firestore.WriteResult, error) {
	degree := strings.ReplaceAll(education.Degree, " ", "-")
	institution := strings.ReplaceAll(education.Institution, " ", "-")

	return u.getEducationCollection(userId).Doc(fmt.Sprintf("%s-%s", degree, institution)).Set(context.TODO(), education.ToMap())
}

func (u User) GetEducation(userId int) ([]schema.Education, error) {
	data, err := u.getEducationCollection(userId).Documents(context.TODO()).GetAll()
	if err != nil {
		return nil, err
	}

	educations := []schema.Education{}

	for _, data := range data {
		educationData := data.Data()

		if educationData["end_date"] == nil {
			educationData["endDate"] = ""
		}

		if educationData["description"] == nil {
			educationData["description"] = ""
		}

		if educationData["start_date"] == nil {
			educationData["startDate"] = ""
		}

		educations = append(educations, schema.Education{
			Degree:      educationData["degree"].(string),
			Institution: educationData["institution"].(string),
			StartDate:   educationData["start_date"].(string),
			EndDate:     educationData["end_date"].(string),
			Description: educationData["description"].(string),
			Hide:        educationData["hide"].(bool),
		})
	}

	return educations, err
}

func (u User) UpdateEducation(userId int, education schema.Education) error {
	doc := u.getEducationCollection(userId).Doc(fmt.Sprintf("%s-%s", education.Degree, education.Institution))
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
