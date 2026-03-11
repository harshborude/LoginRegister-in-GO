package main

import (
	"backend/db"
	"backend/routes"
)

func main() {

	db.ConnectDatabase()

	router := routes.SetupRouter()

	router.Run(":8080")
}