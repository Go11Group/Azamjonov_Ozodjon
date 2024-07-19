package api

import (
	"log"
	"net"

	"auth-service/genprotos/auth_pb"
	"auth-service/internal/config"
	"google.golang.org/grpc"
)

type (
	API struct {
		service auth_pb.AuthServiceServer
	}
)

func New(service auth_pb.AuthServiceServer) *API {
	return &API{
		service: service,
	}
}

func (a *API) RUN(config *config.Config) error {
	listener, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()
	auth_pb.RegisterAuthServiceServer(serverRegisterer, a.service)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
