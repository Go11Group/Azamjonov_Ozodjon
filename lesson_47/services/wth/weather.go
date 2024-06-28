package wth

import (
	"context"
	pb "homeworks/homework_47/genproto/weather"
)

type WeatherRepo struct {
	pb.UnimplementedWeatherServiceServer
}

func (w *WeatherRepo) GetCurrentWeather(ctx context.Context, in *pb.Void) (*pb.WeatherDaily, error) {}
