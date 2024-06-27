package storage

import (
	"database/sql"

	pb "github.com/Azamjonov_Ozodjon/lesson46/genproto/generator/transport" // Import your generated transport proto package
)

type TransportStorage struct {
	db *sql.DB
}

func NewTransportStorage(db *sql.DB) *TransportStorage {
	return &TransportStorage{db: db}
}

func (s *TransportStorage) GetBusSchedule(busNumber string) ([]*pb.BusSchedule, error) {
	query := `SELECT time, destination FROM bus_schedule WHERE bus_number = $1 ORDER BY time`
	rows, err := s.db.Query(query, busNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []*pb.BusSchedule
	for rows.Next() {
		var schedule pb.BusSchedule
		if err := rows.Scan(&schedule.Time, &schedule.Destination); err != nil {
			return nil, err
		}
		schedules = append(schedules, &schedule)
	}
	return schedules, nil
}

func (s *TransportStorage) TrackBusLocation(busNumber string) (string, error) {
	query := `SELECT location FROM bus_location WHERE bus_number = $1 ORDER BY timestamp DESC LIMIT 1`
	row := s.db.QueryRow(query, busNumber)

	var location string
	if err := row.Scan(&location); err != nil {
		return "", err
	}

	return location, nil
}

func (s *TransportStorage) ReportTrafficJam(location, description string) error {
	query := `INSERT INTO traffic_jam (location, description, timestamp) VALUES ($1, $2, NOW())`
	_, err := s.db.Exec(query, location, description)
	return err
}
