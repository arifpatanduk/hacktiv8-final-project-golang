package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	secretKey := os.Getenv("API_SECRET_KEY")
	fmt.Println(secretKey)
}