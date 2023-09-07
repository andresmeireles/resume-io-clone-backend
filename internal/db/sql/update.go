package sql

import (
	"fmt"
	"strings"
)

func (d *Db) Update(model Model, updateValues map[string]interface{}) error {
	defer d.db.Close()
	placeholders := []string{}
	values := []interface{}{}
	placeholder := 1
	for k, v := range updateValues {
		placeholders = append(placeholders, fmt.Sprintf("%s = $%d", k, placeholder))
		values = append(values, v)
		placeholder++
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = %d",
		model.TableName(),
		strings.Join(placeholders, ", "),
		model.GetId(),
	)

	_, err := d.db.Exec(query, values)

	return err
}

func (d *Db) UpdateModel(model Model) error {
	defer d.db.Close()

	fields := ""
	var values []interface{}
	counter := 1

	for k, v := range model.Content() {
		fields += fmt.Sprintf("%s = $%d, ", k, counter)
		values = append(values, v)
		counter++
	}

	fields = strings.Trim(fields, ", ")

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = %d",
		model.TableName(),
		fields,
		*model.GetId(),
	)

	_, err := d.db.Exec(query, values...)

	return err
}
