package sql

import (
	"fmt"
)

func (d *Db) GetOneBy(model Model, identification, value string) error {
	db := d.db

	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", model.TableName(), identification)
	return db.Get(model, query, value)
}
