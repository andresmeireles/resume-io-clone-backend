package user

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
)

func (u User) getProfessionalCollection(userId int) *firestore.CollectionRef {
	return u.getUserCollection(userId).Collection("professional")
}

func (u User) AddProfessionalData(userId int, professional schema.Professional) (*firestore.WriteResult, error) {
	return u.getProfessionalCollection(userId).Doc("professional").Set(context.TODO(), professional.ToMap())
}

func (u User) GetProfessional(userId int) (schema.Professional, error) {
	data, err := u.getProfessionalCollection(userId).Doc("professional").Get(context.TODO())
	if err != nil {
		return schema.Professional{}, err
	}

	return schema.Professional{
		DesiredJob:   data.Data()["desiredJob"].(string),
		Presentation: data.Data()["presentation"].(string),
	}, nil
}
