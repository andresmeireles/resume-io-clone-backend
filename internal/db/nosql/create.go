package nosql

import (
	"context"

	"cloud.google.com/go/firestore"
)

func (n *NoSqlDb) Create(
	ctx context.Context,
	collection, doc string,
	data map[string]interface{},
) (*firestore.WriteResult, error) {
	client := n.Client()

	defer client.Close()

	result, err := client.Collection(collection).Doc(doc).Set(ctx, data)

	if err != nil {
		return nil, err
	}

	return result, nil
}
