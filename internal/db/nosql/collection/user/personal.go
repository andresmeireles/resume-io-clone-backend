package user

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
)

func (u *User) Personal(userId int) *firestore.CollectionRef {
	return u.getUserCollection(userId).Collection("personal")
}

func (u User) GetPersonal(userId int) (schema.Personal, error) {
	data, err := u.getUserCollection(userId).Collection("personal").Doc("personal").Get(context.TODO())
	if err != nil {
		return schema.Personal{}, err
	}
	personalData := data.Data()

	return schema.Personal{
		FirstName: personalData["firstName"].(string),
		LastName:  personalData["lastName"].(string),
		Email:     personalData["email"].(string),
		Age:       int(personalData["age"].(int64)),
		Gender:    personalData["gender"].(string),
		Phone:     personalData["phone"].(string),
		Address:   personalData["address"].(string),
	}, nil
}

func (u *User) AddPersonalData(userId int, personal schema.Personal) (*firestore.WriteResult, error) {
	return u.Personal(userId).Doc("personal").Set(context.TODO(), personal.ToMap())
}
