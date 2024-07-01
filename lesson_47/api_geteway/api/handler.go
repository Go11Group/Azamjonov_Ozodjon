package api

import (
	"homeworks/homework_47/api_geteway/api/handler"
	"homeworks/homework_47/genproto/transport"
	"homeworks/homework_47/genproto/weather"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Handler(conn *grpc.ClientConn) *gin.Engine {

	transport := transport.NewTransportServiceClient(conn)
	weather := weather.NewWeatherServiceClient(conn)
	handler := handler.NewHandlerRepo(transport, weather)

	gin := gin.Default()
	tr := gin.Group("/transport")
	wth := gin.Group("/weather")

	wth.GET("/current", handler.GetCurrentWeather)

	tr.GET("/bus/:id", handler.GetBusSchedule)

	return gin
}
