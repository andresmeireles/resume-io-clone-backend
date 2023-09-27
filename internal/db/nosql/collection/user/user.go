package user

import (
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/andresmeireles/resume/internal/db/nosql"
)

const (
	skill      = "skill"
	personal   = "personal"
	experience = "experience"
)

type User struct {
	userCollection *firestore.CollectionRef
}

func Instance() *User {
	client := nosql.GetNoSqlClient()
	return &User{userCollection: client.Client().Collection("users")}
}

func (u *User) getUserCollection(userId int) *firestore.DocumentRef {
	return u.userCollection.Doc(strconv.Itoa(userId))
}
