package utils

import "github.com/joho/godotenv"

// GetEnv retrieves the value of an environment variable by its key.
//
// Parameters:
// - key: the key of the environment variable to retrieve.
//
// Return:
// - interface{}: the value of the environment variable as string, or nil if it doesn't exist.
func GetEnv(key string) interface{} {
	env, exists := env()[key]
	if exists {
		return env
	}
	return nil
}

func GetEnvAsString(key string) string {
	env := GetEnv(key)
	if env != nil {
		return env.(string)
	}
	return ""
}

func env() map[string]string {
	env, err := godotenv.Read()
	if err != nil {
		panic(err)
	}

	return env
}
