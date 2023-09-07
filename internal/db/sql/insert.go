package sql

import (
	"fmt"
	"strings"
)

func (d *Db) Insert(model Model) error {
	db := d.db

	defer db.Close()

	keys := []string{}
	value := []interface{}{}

	for k, v := range model.InsertFields() {
		keys = append(keys, k)
		value = append(value, v)
	}

	tableName := model.TableName()
	var placeHolders string
	for _, v := range keys {
		placeHolders += ":" + v + ", "
	}
	placeHolders = strings.Trim(placeHolders, ", ")

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES(%s)",
		tableName,
		strings.Trim(strings.Join(keys, ","), ","),
		placeHolders,
	)

	result, err := db.NamedExec(query, model)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}
