package sql

import (
	"fmt"

	"github.com/andresmeireles/resume/utils/env"
	"github.com/andresmeireles/resume/utils/path"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Up() {
	m := getMigrate()

	m.Up()
}

func Down() {
	m := getMigrate()

	m.Down()
}

func getMigrate() *migrate.Migrate {

	db := GetDB().Instance()

	defer db.Close()

	migrationsFolder := path.GetRootDir() + "/db/sql/migration"

	dbName := env.GetEnvAsString("DB_NAME")

	if dbName == "" {
		panic("DB_NAME is not set")
	}

	driver, e := postgres.WithInstance(db.DB, &postgres.Config{})
	if e != nil {
		panic(e)
	}

	fmt.Println(migrationsFolder)

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsFolder,
		dbName,
		driver,
	)

	if err != nil {
		panic(err)
	}

	return m
}
