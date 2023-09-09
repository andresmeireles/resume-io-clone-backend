package nosql

import (
	"context"

	"cloud.google.com/go/firestore"
)

func (n *NoSqlDb) Update(ctx context.Context, collection, doc string, data map[string]interface{}) (*firestore.WriteResult, error) {
	client := n.Client()

	defer client.Close()

	updateData := castDataToUpdate(data)

	return client.Collection(collection).Doc(doc).Update(ctx, updateData)
}

func castDataToUpdate(data map[string]interface{}) []firestore.Update {
	var updates []firestore.Update
	for k, v := range data {
		updates = append(updates, firestore.Update{
			Path:  k,
			Value: v,
		})
	}
	return updates
}
