package main

import (
	"go-mygram/config"
	router "go-mygram/routers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {

	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	config.StartDB()

	port := ":" + os.Getenv("APP_PORT")
	router.StartApp().Run(port)
}