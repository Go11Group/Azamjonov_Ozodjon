package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	pb "github.com/Azamjonov_Ozodjon/lesson46/genproto/generator/weather" // Import your generated weather proto package
	"github.com/Azamjonov_Ozodjon/lesson46/storage"
)

type WeatherService struct {
	storage *storage.WeatherStorage
}

func NewWeatherService(db *storage.DB) *WeatherService {
	return &WeatherService{
		storage: storage.NewWeatherStorage(db),
	}
}

func (s *WeatherService) GetCurrentWeather(ctx context.Context, req *pb.CurrentWeatherRequest) (*pb.CurrentWeatherResponse, error) {
	weather, err := s.storage.GetCurrentWeather(ctx, req.Location)
	if err != nil {
		log.Printf("Failed to get current weather: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get current weather: %v", err)
	}

	response := &pb.CurrentWeatherResponse{
		Temperature: weather.Temperature,
		Humidity:    weather.Humidity,
		WindSpeed:   weather.WindSpeed,
	}

	return response, nil
}

func (s *WeatherService) ReportWeatherCondition(ctx context.Context, req *pb.WeatherConditionRequest) (*pb.WeatherConditionResponse, error) {
	err := s.storage.ReportWeatherCondition(ctx, req.Weather)
	if err != nil {
		log.Printf("Failed to report weather condition: %v", err)
		return &pb.WeatherConditionResponse{Success: false}, status.Errorf(codes.Internal, "failed to report weather condition: %v", err)
	}

	return &pb.WeatherConditionResponse{Success: true}, nil
}
