package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type config struct{
	Port string `json:"port"`
}

var AppConfig config

func SetConfig(){
	err := godotenv.Load("./config/.env")

	if err != nil {
		log.Fatal("env file does not exist")
	}

	AppConfig.Port = os.Getenv("Port")
}