package nosql

import "cloud.google.com/go/firestore"

func (n NoSqlDb) Client() *firestore.Client {
	return n.client
}
