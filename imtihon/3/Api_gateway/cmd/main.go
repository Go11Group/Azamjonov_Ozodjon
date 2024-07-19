package main

import (
	"log"

	"github.com/imtihon/3/Api_gateway/api"
	hendler "github.com/imtihon/3/Api_gateway/api/handler"
	itempb "github.com/imtihon/3/Api_gateway/generated/generated/items_service"
	authpb "github.com/imtihon/3/Api_gateway/generated/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the auth service
	authConn, err := grpc.NewClient(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to auth service: %v", err)
	}
	defer authConn.Close()

	// Connect to the item service
	itemConn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to item service: %v", err)
	}
	defer itemConn.Close()

	// Create service clients
	authClient := authpb.NewAuthServiceClient(authConn)
	itemClient := itempb.NewItemsServiceClient(itemConn)

	// Create the handler with the clients
	h := hendler.Handler{
		Auth:  authClient,
		Items: itemClient,
	}

	// Initialize the router with the handler
	r := api.NewRouter(h)

	// Start the server
	if err := r.Run(":50053"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
