package main

import (
	"fmt"
	postgres2 "github.com/Azamjonov_Ozodjon/lesson_45/storage/postgres"
	"log"
	"net"

	pb "github.com/Azamjonov_Ozodjon/lesson_45/generator/library"
	"github.com/Azamjonov_Ozodjon/lesson_45/service"
	"github.com/Azamjonov_Ozodjon/lesson_45/storage"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres2.Connection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	bookStorage := storage.NewBookStorage(db)
	s := service.NewServer(bookStorage)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLibraryServiceServer(grpcServer, s)

	fmt.Println("Server is running on port :50051")
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
