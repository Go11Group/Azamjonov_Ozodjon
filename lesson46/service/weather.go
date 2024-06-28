package service

import (
	"context"
	"database/sql"
	"log"
	"time"

	pb "github.com/Azamjonov_Ozodjon/lesson46/genproto/generator/weather"
	"github.com/Azamjonov_Ozodjon/lesson46/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WeatherService struct {
	pb.UnimplementedWeatherServiceServer // Embed the unimplemented server
	storage                              *storage.WeatherStorage
}

func NewWeatherService(db *sql.DB) *WeatherService {
	return &WeatherService{
		storage: storage.NewWeatherStorage(db),
	}
}

func (s *WeatherService) GetCurrentWeather(ctx context.Context, req *pb.CurrentWeatherRequest) (*pb.CurrentWeatherResponse, error) {
	weather, err := s.storage.GetCurrentWeather(req.Location)
	if err != nil {
		log.Printf("Failed to get current weather: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get current weather: %v", err)
	}

	return weather, nil
}

func (s *WeatherService) GetWeatherForecast(ctx context.Context, req *pb.WeatherForecastRequest) (*pb.WeatherForecastResponse, error) {
	// Implementation to get weather forecast (you need to add the database queries and logic)
	// This is a placeholder implementation
	forecast := &pb.WeatherForecastResponse{
		Forecasts: []*pb.WeatherForecast{
			{
				Date:        time.Now().Add(24 * time.Hour).Format("2006-01-02"),
				Temperature: 25.0,
				Humidity:    60.0,
				WindSpeed:   5.0,
			},
		},
	}
	return forecast, nil
}

func (s *WeatherService) ReportWeatherCondition(ctx context.Context, req *pb.ReportWeatherConditionRequest) (*pb.ReportWeatherConditionResponse, error) {
	err := s.storage.ReportWeatherCondition(req.Location, req.Condition)
	if err != nil {
		log.Printf("Failed to report weather condition: %v", err)
		return &pb.ReportWeatherConditionResponse{Success: false}, status.Errorf(codes.Internal, "failed to report weather condition: %v", err)
	}

	return &pb.ReportWeatherConditionResponse{Success: true}, nil
}
