package nosql

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/andresmeireles/resume/utils"
)

func connect(ctx context.Context, operation func(ctx context.Context, client *firestore.Client) interface{}) interface{} {
	conf := &firebase.Config{ProjectID: utils.GetEnvAsString("FIRESTORE_PROJECT_ID")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	return operation(ctx, client)
}

func Run(ctx context.Context, operation func(ctx context.Context, client *firestore.Client) interface{}) interface{} {
	return connect(ctx, operation)
}
