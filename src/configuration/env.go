package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	PORT      string
	DB_URL    string
	WS_URL    string
	WS_ORIGIN string
}

func Env() *env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	enviroment := new(env)
	enviroment.PORT = os.Getenv("PORT")
	enviroment.DB_URL = os.Getenv("DB_URL")
	enviroment.WS_ORIGIN = os.Getenv("WS_ORIGIN")
	enviroment.WS_URL = os.Getenv("WS_URL")

	return enviroment
}
