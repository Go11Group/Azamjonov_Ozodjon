package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Azamjonov_Ozodjon/lesson44/proto/translator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewTranslatorServerClient(conn)

	req := pb.Request{Words: []string{"apple", "book", "house", "car", "tree", "water", "fire", "sky", "earth"}}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Translate(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)
}
