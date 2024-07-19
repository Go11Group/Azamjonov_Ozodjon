package main

import (
	"github.com/imtihon/3/Item_Service/logs"
	"log"
	"net"

	pb "github.com/imtihon/3/Item_Service/generated/generated/items_service"
	"github.com/imtihon/3/Item_Service/servie"
	"github.com/imtihon/3/Item_Service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {

	// Initialize logger
	logs.InitLogger()
	logs.Logger.Info("Starting the server")

	// Connect to PostgreSQL database
	db, err := postgres.Conn()
	if err != nil {
		logs.Logger.Error("Failed to connect to Data Base", "error", err.Error())
		panic(err)
	}
	defer db.Close()

	// Listen on the specified port
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logs.Logger.Error("Failed to listen", "error", err.Error())
		panic(err)
	}

	// Create repositories
	s := service.NewItemService(*postgres.NewItemRepo(db))
	server := grpc.NewServer()
	pb.RegisterItemsServiceServer(server, s)

	// Log server startup
	logs.Logger.Info("Server is Running", "PORT", "50051")
	log.Println("Server is running on ", listener.Addr())

	// Start serving
	if err = server.Serve(listener); err != nil {
		logs.Logger.Error("Failed to run the server", "error", err.Error())
		panic(err)
	}
}
