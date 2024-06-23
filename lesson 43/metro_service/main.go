package main

import (
	"atto/api"
	"atto/storage/postgres"
	"log"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := api.Routes(db)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
