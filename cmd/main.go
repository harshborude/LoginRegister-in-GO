package main

import (
	"backend/db"
	"backend/routes"
	"backend/utils"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// initialize JWT secrets AFTER env loading
	utils.InitJWT()

	db.ConnectDatabase()

	router := routes.SetupRouter()

	router.Run(":8080")
}