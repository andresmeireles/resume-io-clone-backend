package resume

import (
	"context"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/andresmeireles/resume/internal/db/nosql"
	"github.com/andresmeireles/resume/internal/db/nosql/schema"
	"github.com/andresmeireles/resume/internal/db/sql/model"
)

const (
	beginner = "beginner"
)

func Add(client nosql.NoSqlDbInterface, user model.User, skill schema.Skill) (*firestore.WriteResult, error) {
	return client.Create(context.Background(), "users", strconv.Itoa(int(*user.GetId())), skill.ToMap())
}

func Update() {}

func Remove() {}
