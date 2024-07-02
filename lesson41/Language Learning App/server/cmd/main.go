package main

import (
	"language_learning_app/api/handler"
	"language_learning_app/api/router"
	"language_learning_app/storage/postgres"
)

func main() {
	db, err := postgres.CreateDB()
	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(db)
	server := router.CreateServer(h)
	server.ListenAndServe()
}
