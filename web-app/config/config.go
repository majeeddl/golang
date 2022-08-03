package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct{
	PORT string `json:"port"`
}

var AppConfig config

func SetConfig(){
	err := godotenv.Load("./config/.env")


	if err != nil {
		log.Fatal("env file does not exist")
	}
	
	AppConfig.PORT = os.Getenv("PORT")
}