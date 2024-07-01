package postgers

import (
	"database/sql"
	_"github.com/lib/pq"
	"fmt"
)

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "dars_5"
	password = "123321"
)

func Conn() (*sql.DB, error){
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
	host, port, user, dbname, password)

	db, err := sql.Open("postgres", conn)
	if err != nil{
		return nil, err
	}
	return db, err
}