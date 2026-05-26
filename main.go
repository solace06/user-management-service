package main

import (
	"log"

	"github.com/solace06/user-management-service/api"
	"github.com/solace06/user-management-service/database"
)

func main() {

	//database connection
	_, err := database.NewDB()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	//router setup
	r := api.NewRouter()

	//start server
	r.Run(":8080")
}
