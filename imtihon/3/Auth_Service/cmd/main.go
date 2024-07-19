package main

import (
	"log"
	"log/slog"
	"os"

	"auth-service/api"
	"auth-service/internal/config"
	"auth-service/internal/service"
	"auth-service/internal/storage"
)

func main() {
	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(logFile, nil))

	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	storage, err := storage.New(configs, logger)
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(service.New(*storage, logger))

	log.Fatal(api.RUN(configs))
}
