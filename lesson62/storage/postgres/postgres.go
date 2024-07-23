package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	conn := `host=localhost port=5432 user=postgres dbname=ok 
	password=BEKJONS sslmode=disable`

	db, err := sql.Open("postgres", conn)
	return db, err
}
