package main

import (
	"backend/db"
	"backend/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	db.ConnectDatabase()

	router := routes.SetupRouter()

	router.Run(":8080")
}