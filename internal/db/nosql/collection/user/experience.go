package user

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
)

func (u *User) GetExperience(userId int) *firestore.CollectionRef {
	return u.getUserCollection(userId).Collection(experience)
}

func (u *User) AddExperience(userId int, experience schema.Experience) (*firestore.WriteResult, error) {
	return u.GetSkills(userId).Doc(fmt.Sprintf("%s-%s", experience.Job, experience.Company)).Set(context.TODO(), experience.ToMap())
}

func (u *User) UpdateExperience(userId int, experience schema.Experience) error {
	doc := u.GetSkills(userId).Doc(fmt.Sprintf("%s-%s", experience.Job, experience.Company))
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
