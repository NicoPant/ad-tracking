package main

import (
	"ad/config"
	"ad/db"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()

	err = db.InitMongo(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
