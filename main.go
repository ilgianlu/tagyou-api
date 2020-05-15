package main

import (
	"log"
	"os"

	"github.com/ilgianlu/tagyou-api/api"
	dotenv "github.com/joho/godotenv"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
		return
	}

	var api api.ApiServer
	api.Start(os.Getenv("LISTEN_PORT"))
}

func loadEnv() error {
	env := os.Getenv("TAGYOU_ENV")
	if "" == env {
		env = "default"
	}
	return dotenv.Load(".env." + env + ".local")
}
