package sql

import (
	"fmt"

	"github.com/andresmeireles/resume/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	host := utils.GetEnv("DB_HOST")
	user := utils.GetEnv("DB_USER")
	password := utils.GetEnv("DB_PASSWORD")
	name := utils.GetEnv("DB_NAME")
	port := utils.GetEnv("DB_PORT")
	timezone := utils.GetEnv("DB_TIMEZONE")
	if timezone == "" {
		timezone = "America/Belem"
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		host,
		user,
		password,
		name,
		port,
		timezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
