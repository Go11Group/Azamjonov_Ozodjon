package main

import (
	"fmt"
	"log"
	"net"

	pb_transport "github.com/Azamjonov_Ozodjon/lesson46/genproto/generator/transport"
	pb_weather "github.com/Azamjonov_Ozodjon/lesson46/genproto/generator/weather"
	"github.com/Azamjonov_Ozodjon/lesson46/service"
	"github.com/Azamjonov_Ozodjon/lesson46/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.Connection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	weatherService := service.NewWeatherService(db)
	transportService := service.NewTransportService(db)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb_weather.RegisterWeatherServiceServer(grpcServer, weatherService)
	pb_transport.RegisterTransportServiceServer(grpcServer, transportService)

	fmt.Println("Server is running on port :50051")
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
