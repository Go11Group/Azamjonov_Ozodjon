package storage

import (
	"database/sql"
	"time"

	pb "github.com/Azamjonov_Ozodjon/lesson46/genproto/generator/weather" // Import your generated weather proto package
)

type WeatherStorage struct {
	db *sql.DB
}

func NewWeatherStorage(db *sql.DB) *WeatherStorage {
	return &WeatherStorage{db: db}
}

func (s *WeatherStorage) GetCurrentWeather(location string) (*pb.Weather, error) {
	query := `SELECT temperature, humidity, wind_speed FROM weather WHERE location = $1 ORDER BY timestamp DESC LIMIT 1`
	row := s.db.QueryRow(query, location)

	var weather pb.Weather
	if err := row.Scan(&weather.Temperature, &weather.Humidity, &weather.WindSpeed); err != nil {
		return nil, err
	}

	return &weather, nil
}

func (s *WeatherStorage) ReportWeatherCondition(weather *pb.WeatherCondition) error {
	query := `INSERT INTO weather (location, temperature, humidity, wind_speed, timestamp) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.Exec(query, weather.Location, weather.Temperature, weather.Humidity, weather.WindSpeed, time.Now())
	return err
}
