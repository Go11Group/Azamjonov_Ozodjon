package wth

import (
	"context"
	pb "homeworks/homework_47/genproto/weather"
)

func (w *WeatherRepo) GetCurrentWeather(ctx context.Context, in *pb.Void) (*pb.WeatherDaily, error) {
	weather := pb.WeatherCondition{Temperature: 30, Humidity: "15%", WindSpeed: 13, Condition: "rainy"}
	return &pb.WeatherDaily{Date: "2022-06-16", Location: "Dubai", Weather: &weather}, nil
}
