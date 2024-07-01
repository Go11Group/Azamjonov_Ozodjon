package handler

import (
	"homeworks/homework_47/genproto/transport"
	"homeworks/homework_47/genproto/weather"
)

type HandlerRepo struct {
	Transport transport.TransportServiceClient
	Weather weather.WeatherServiceClient
}

func NewHandlerRepo(t transport.TransportServiceClient, w weather.WeatherServiceClient) *HandlerRepo{
	return &HandlerRepo{t, w}
}