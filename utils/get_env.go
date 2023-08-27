package utils

import "github.com/joho/godotenv"

func GetEnv(key string) string {
	return env()[key]
}

func env() map[string]string {
	env, err := godotenv.Read()
	if err != nil {
		panic(err)
	}

	return env
}
