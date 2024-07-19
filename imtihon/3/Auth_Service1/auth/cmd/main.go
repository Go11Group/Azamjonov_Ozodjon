package main

import (
	"auth_service/api"
	"auth_service/api/handler"
	"auth_service/logs"
	"auth_service/service"
	"auth_service/storage/postgres"
	"fmt"
	"net/http"
)

func main() {
	logger := logs.InitLogger()
	// Initialize database connection
	db, err := postgres.Conn()
	if err != nil {
		logger.Error("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize service layer with repository, logger, and ecoPoints
	userAuthRepo := postgres.NewUserAuthRepo(db)
	ecoPointsRepo := postgres.NewEcoPointsRepo(db)
	userService := service.NewUserService(userAuthRepo, logger, ecoPointsRepo)

	// Initialize handler with logger and service layer
	authHandler := handler.NewHandler(logger, userService)

	// Initialize API router
	router := api.Router(authHandler, db)

	// Start the server
	port := ":50052"
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		logger.Error("Failed to start server: %v", err)
	}
}
