package storage

import (
	"database/sql"
	"fmt"
)

type StorageRepo struct {
	DB *sql.DB
}

type trow struct {
	From string
	To   string
}

type jrow struct{ IsTraficJam bool }

func NewStorageRepo(db *sql.DB) StorageRepo {
	return StorageRepo{db}
}

func (s *StorageRepo) GetScheduleByNumber(number int) (string, error) {
	r := trow{}
	row := s.DB.QueryRow("select * from bus where number = $1", number)
	row.Scan(&r.From, &r.To)
	str := fmt.Sprint(r)
	return str, row.Err()
}

func (s *StorageRepo) IsTraficJam(street string) (bool, error) {
	r := jrow{}
	row := s.DB.QueryRow("select * from trafic_jam where street = $1", street)
	row.Scan(&r.IsTraficJam)
	return r.IsTraficJam, row.Err()
}
