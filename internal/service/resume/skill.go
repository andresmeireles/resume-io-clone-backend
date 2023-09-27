package resume

import (
	"cloud.google.com/go/firestore"
	collection "github.com/andresmeireles/resume/internal/db/nosql/collection/user"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
	"github.com/andresmeireles/resume/internal/db/sql/model"
)

func AddSkill(userCollection *collection.User, user model.User, skill schema.Skill) (*firestore.WriteResult, error) {
	return userCollection.AddSkill(int(*user.Id), skill)
}

func UpdateSkill(userCollection *collection.User, user model.User, skill schema.Skill) error {
	return userCollection.UpdateSkill(int(*user.Id), skill)
}

func RemoveSkill() {}
