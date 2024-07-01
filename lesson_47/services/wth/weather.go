package wth

import (
	pb "homeworks/homework_47/genproto/weather"
)

type WeatherRepo struct {
	pb.UnimplementedWeatherServiceServer
}
