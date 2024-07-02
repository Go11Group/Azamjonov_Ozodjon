package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Azamjonov_Ozodjon/lesson46/protos/genproto/protos"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeWeatherServiceServer
}

func (s *server) GetCurrentWeather(ctx context.Context, in *pb.CurrentWeatherRequest) (*pb.CurrentWeatherResponse, error) {
	log.Printf("Received GetCurrentWeather request for location: %s", in.Location)

	return &pb.CurrentWeatherResponse{
		Location: in.Location,
		Obuhavo:  "Rainy",
		Harorat:  36.2,
	}, nil
}

func (s *server) GetWeatherForecast(ctx context.Context, in *pb.ForecastRequest) (*pb.ForecastResponse, error) {
	log.Printf("Received GetWeatherForecast request for location: %s, days: %d", in.Location, in.Kunlar)

	forecasts := []*pb.Forecast{
		{Date: "2024-06-28", Weather: "Sunny"},
		{Date: "2024-06-29", Weather: "Cloudy"},
		{Date: "2024-06-30", Weather: "Rainy"},
	}
	return &pb.ForecastResponse{
		Location: in.Location,
		Forecast: forecasts,
	}, nil
}

func (s *server) ReportWeatherCondition(ctx context.Context, in *pb.ConditionReport) (*pb.ConditionResponse, error) {
	log.Printf("Received ReportWeatherCondition request for location: %s, weather: %s, temperature: %d", in.Location, in.Weather, in.Temperature)

	return &pb.ConditionResponse{
		Success: true,
		Message: "Weather condition reported successfully.",
	}, nil
}

func (s *server) mustEmbedUnimplementedWeatherServiceServer() {}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, &server{})
	log.Println("Starting gRPC server on port :50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
