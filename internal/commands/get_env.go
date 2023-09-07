package commands

import "os"

func getEnv() *os.File {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.Open(currentDir + "/../.env")

	if err != nil {
		panic(err)
	}

	return file
}
