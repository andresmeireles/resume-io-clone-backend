package nosql

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/andresmeireles/resume/utils/env"
)

type NoSqlDbInterface interface {
	Create(ctx context.Context, collection, doc string, data map[string]interface{}) (*firestore.WriteResult, error)
	Update(ctx context.Context, collection, doc string, data map[string]interface{}) (*firestore.WriteResult, error)
	Client() *firestore.Client
}

type NoSqlDb struct {
	client *firestore.Client
}

func GetNoSqlClient() NoSqlDb {
	return NoSqlDb{client: client()}
}

func client() *firestore.Client {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: env.GetEnvAsString("FIRESTORE_PROJECT_ID")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
