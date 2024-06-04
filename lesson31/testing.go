package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-faker/faker/v4"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "postgres"
	dbname   = "ok"
	password = "BEKJONS"
	port     = 5432
)

type z struct{}

func main() {
	dbCon := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		host, user, dbname, password, port)

	db, err := sql.Open("postgres", dbCon)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	v := map[int]z{}

	var a int
	q := z{}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	for i := 0; i < 600; i++ {
		// model := faker.Name()
		go func() {
			year := faker.FirstName()
			// num := faker.SetRandomNumberBoundaries(1000, 2000)
			// color := faker.Word()
			// owner := faker.Name()
			t := time.Now()
			err = db.QueryRow("select count(1) from product where name = $1",
				year).Scan(&a)
			if err != nil {
				fmt.Println(err)
				return
			}
			v[i] = q
			fmt.Println(i, a, time.Now().Sub(t))
		}()
		if i%1000 == 0 {
			fmt.Println(i)
		}
	}
	time.Sleep(5 * time.Second)
	fmt.Println("rtyui", len(v))
}
