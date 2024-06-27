package main

import (
	"github.com/Azamjonov_Ozodjon/lesson46/storage"
	"github.com/Azamjonov_Ozodjon/lesson46/storage/postgres"
	"log"
)

func main() {
	db, err := postgres.Connection()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	weatherStorage := storage.NewWeatherStorage(db)
	transportStorage := storage.NewTransportStorage(db)

	// Example calls to storage methods
	weather, err := weatherStorage.GetCurrentWeather("New York")
	if err != nil {
		log.Fatalf("Failed to get current weather: %v", err)
	}
	log.Printf("Current Weather: %+v\n", weather)

	schedules, err := transportStorage.GetBusSchedule("42")
	if err != nil {
		log.Fatalf("Failed to get bus schedule: %v", err)
	}
	log.Printf("Bus Schedules: %+v\n", schedules)
}
