package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "BEKJONS"
	dbname   = "ok"
)

func Connection() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
		return nil, err
	}

	return db, nil
}
