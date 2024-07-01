package main

import (
	"homeworks/homework_47/api_geteway/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	router := api.Handler(conn)
	panic(router.Run(":8080"))
}
