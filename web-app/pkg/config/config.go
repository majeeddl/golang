package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	PORT string `json:"port"`
}

var AppConfig config

func SetConfig() {
	err := godotenv.Load("./pkg/config/.env")

	if err != nil {
		log.Fatal("env file does not exist")
	}

	AppConfig.PORT = os.Getenv("PORT")
}
