package main

import (
	"homeworks/homework_47/genproto/transport"
	"homeworks/homework_47/genproto/weather"
	"homeworks/homework_47/services/tr"
	"homeworks/homework_47/services/wth"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listner, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}
	defer listner.Close()

	server := grpc.NewServer()

	tr := tr.TransportRepo{}
	wth := wth.WeatherRepo{}

	transport.RegisterTransportServiceServer(server, &tr)
	weather.RegisterWeatherServiceServer(server, &wth)

	log.Println("server is running on :8090...")

	panic(server.Serve(listner))
}
