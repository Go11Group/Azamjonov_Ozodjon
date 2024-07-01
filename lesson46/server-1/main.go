package main

import (
	pb "github.com/Azamjonov_Ozodjon/lesson46/generate/weather"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listner, err := net.Listen("tcp", ":5051")
	if err != nil {
		panic(err)
	}

	s := &server{}
	grpc := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, grpc)

	err = grpc.Server(listner)
	if err != nil {
		panic(err)
	}
}
