package postgres

import (
	"database/sql"
	"time"

	"atto/models"
	"github.com/google/uuid"
)

type StationRepo struct {
	Db *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{Db: db}
}

func (s *StationRepo) Create(station *models.CreateStation) error {
	_, err := s.Db.Exec("INSERT INTO station (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)",
		uuid.NewString(), station.Name, time.Now(), time.Now())
	return err
}

func (s *StationRepo) GetById(id string) (*models.Station, error) {
	var station models.Station
	err := s.Db.QueryRow("SELECT id, name, created_at, updated_at, deleted_at FROM station WHERE id = $1", id).
		Scan(&station.Id, &station.Name, &station.CreatedAt, &station.UpdatedAt, &station.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &station, nil
}

func (s *StationRepo) Update(station *models.Station) error {
	_, err := s.Db.Exec("UPDATE station SET name = $2, updated_at = $3 WHERE id = $1",
		station.Id, station.Name, time.Now())
	return err
}

func (s *StationRepo) Delete(id string) error {
	_, err := s.Db.Exec("UPDATE station SET deleted_at = $2 WHERE id = $1",
		id, time.Now())
	return err
}

func (s *StationRepo) GetAll() ([]models.Station, error) {
	rows, err := s.Db.Query("SELECT id, name, created_at, updated_at, deleted_at FROM station WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stations []models.Station
	for rows.Next() {
		var station models.Station
		if err := rows.Scan(&station.Id, &station.Name, &station.CreatedAt, &station.UpdatedAt, &station.DeletedAt); err != nil {
			return nil, err
		}
		stations = append(stations, station)
	}
	return stations, nil
}
