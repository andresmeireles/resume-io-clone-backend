package user

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
)

func (u User) getSocialCollection(userId int) *firestore.CollectionRef {
	return u.getUserCollection(userId).Collection("social")
}

func (u User) AddSocial(userId int, social schema.Social) (*firestore.WriteResult, error) {
	return u.getSocialCollection(userId).Doc(social.Name).Set(context.TODO(), social.ToMap())
}

func (u User) GetSocial(userId int) ([]schema.Social, error) {
	data, err := u.getSocialCollection(userId).Documents(context.TODO()).GetAll()
	if err != nil {
		return nil, err
	}
	socials := []schema.Social{}

	for _, data := range data {
		socialData := data.Data()
		socials = append(socials, schema.Social{
			Name: socialData["name"].(string),
			Link: socialData["link"].(string),
			Hide: socialData["hide"].(bool),
		})
	}

	return socials, err
}

func (u User) UpdateSocial(userId int, social schema.Social) error {
	doc := u.getSocialCollection(userId).Doc(social.Name)
	getDoc, err := doc.Get(context.TODO())
	if err != nil {
		return err
	}
	if getDoc.Exists() {
		doc.Set(context.TODO(), social.ToMap())
		return nil
	}

	return errors.New("Social not found")
}
