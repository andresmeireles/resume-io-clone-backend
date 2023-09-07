package sql

import "github.com/jmoiron/sqlx"

type DBInterface interface {
	Update(model Model, updateValues map[string]interface{}) error
	Insert(model Model) error
	Instance() *sqlx.DB
	GetOneBy(model Model, identification, value string) error
	UpdateModel(model Model) error
}
