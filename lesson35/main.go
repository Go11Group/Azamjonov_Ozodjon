package main

import (
	"github.com/Azamjonov_Ozodjon/lesson35/handler"
	"github.com/Azamjonov_Ozodjon/lesson35/storage/postgres"
	"log"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	usersRepo := postgres.UsersRepo{Db: db}
	problemsRepo := postgres.ProblemsRepo{Db: db}
	sproblemsRepo := postgres.SolvedProblemsRepo{Db: db}
	server := handler.NewHandler(usersRepo, problemsRepo, sproblemsRepo)
	log.Fatal(server.ListenAndServe())
}
