package sql

import (
	"fmt"

	"github.com/andresmeireles/resume/utils"
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

	migrationsFolder := utils.GetRootDir() + "/db/sql/migration"

	dbName := utils.GetEnvAsString("DB_NAME")

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
