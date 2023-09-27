package user

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
)

func (u *User) Personal(userId int) *firestore.CollectionRef {
	return u.getUserCollection(userId).Collection("personal")
}

func (u *User) AddPersonalData(userId int, personal schema.Personal) (*firestore.WriteResult, error) {
	return u.Personal(userId).Doc("personal").Set(context.TODO(), personal.ToMap())
}
