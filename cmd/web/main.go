package main

import (
	"os"

	"github.com/andresmeireles/resume/web"
	"github.com/joho/godotenv"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	e := godotenv.Load(dir + "/.env")
	if e != nil {
		panic(e)
	}
	web.Run()
}
