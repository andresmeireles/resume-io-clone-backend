package sql

import (
	"fmt"

	"github.com/andresmeireles/resume/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Db struct {
	db *sqlx.DB
}

func GetDB() *Db {
	return &Db{db: connect()}
}

func (d Db) Instance() *sqlx.DB {
	return d.db
}

func connect() *sqlx.DB {
	user := utils.GetEnvAsString("DB_USER")
	password := utils.GetEnvAsString("DB_PASSWORD")
	dbName := utils.GetEnvAsString("DB_NAME")
	host := utils.GetEnvAsString("DB_HOST")
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		user, password, dbName, host,
	)
	connect, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return connect
}
