package main

import (
	"database/sql"
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:BEKJONS@localhost:5432/ok?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for i := 0; i < 1000000; i++ {
		_, err = db.Exec("insert into product (id,name, category, cost) values ($1, $2, $3, $4)",
			uuid.NewString(), faker.FirstName(), faker.LastName(), 4234)
		if err != nil {
			fmt.Println(err)
		}
		if i%10000 == 0 {
			fmt.Println(i)
		}
	}
}
