package resume

import (
	"cloud.google.com/go/firestore"
	collection "github.com/andresmeireles/resume/internal/db/nosql/collection/user"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
	"github.com/andresmeireles/resume/internal/db/sql/model"
)

func AddPersonal(userCollection *collection.User, user model.User, personal schema.Personal) (*firestore.WriteResult, error) {
	return userCollection.AddPersonalData(int(*user.Id), personal)
}

func UpdatePersonal() {}

func RemovePersonal() {}
