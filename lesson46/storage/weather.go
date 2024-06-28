package storage

import (
	"database/sql"
	"time"

	pb "github.com/Azamjonov_Ozodjon/lesson46/genproto/generator/weather"
)

type WeatherStorage struct {
	db *sql.DB
}

func NewWeatherStorage(db *sql.DB) *WeatherStorage {
	return &WeatherStorage{db: db}
}

func (s *WeatherStorage) GetCurrentWeather(location string) (*pb.CurrentWeatherResponse, error) {
	query := `SELECT temperature, humidity, wind_speed FROM weather WHERE location = $1 ORDER BY timestamp DESC LIMIT 1`
	row := s.db.QueryRow(query, location)

	var weather pb.CurrentWeatherResponse
	if err := row.Scan(&weather.Temperature, &weather.Humidity, &weather.WindSpeed); err != nil {
		return nil, err
	}

	return &weather, nil
}

func (s *WeatherStorage) ReportWeatherCondition(location, condition string) error {
	query := `INSERT INTO weather (location, temperature, humidity, wind_speed, timestamp) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.Exec(query, location, condition, time.Now())
	return err
}
