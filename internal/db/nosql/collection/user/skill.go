package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
)

func (u *User) getSkills(userId int) *firestore.CollectionRef {
	return u.getUserCollection(userId).Collection(skill)
}

func (u User) GetAllSkills(userId int) ([]schema.Skill, error) {
	allSkills := []schema.Skill{}
	skills, err := u.getSkills(userId).Documents(context.TODO()).GetAll()
	if err != nil {
		return nil, err
	}
	for _, skill := range skills {
		skillData := skill.Data()
		allSkills = append(allSkills, schema.Skill{
			Name:  skillData["name"].(string),
			Level: skillData["level"].(string),
			Hide:  skillData["hide"].(bool),
		})
	}

	return allSkills, nil
}

func (u *User) AddSkill(userId int, skill schema.Skill) (*firestore.WriteResult, error) {
	return u.getSkills(userId).Doc(skill.Name).Set(context.TODO(), skill.ToMap())
}

func (u *User) UpdateSkill(userId int, skill schema.Skill) error {
	doc := u.getSkills(userId).Doc(strings.Trim(skill.Name, " "))
	getDoc, err := doc.Get(context.TODO())
	if err != nil {
		return errors.New(fmt.Sprintf("Skill %s not found", skill.Name))
	}
	if getDoc.Exists() {
		doc.Set(context.TODO(), skill.ToMap())
		return nil
	}

	return errors.New("Skill not found")
}
