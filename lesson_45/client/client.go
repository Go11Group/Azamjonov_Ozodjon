package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Azamjonov_Ozodjon/lesson_45/generator/library"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewLibraryServiceClient(conn)

	// Add a book
	addBookReq := &pb.AddBookRequest{
		Title:         "1984",
		Author:        "George Orwell",
		YearPublished: 1949,
	}
	addBookRes, err := client.AddBook(context.Background(), addBookReq)
	if err != nil {
		log.Fatalf("AddBook failed: %v", err)
	}
	fmt.Printf("Added Book ID: %s\n", addBookRes.BookId)

	// Search for a book
	searchBookReq := &pb.SearchBookRequest{Query: "George Orwell"}
	searchBookRes, err := client.SearchBook(context.Background(), searchBookReq)
	if err != nil {
		log.Fatalf("SearchBook failed: %v", err)
	}
	fmt.Println("Search Results:")
	for _, book := range searchBookRes.Books {
		fmt.Printf("Book ID: %s, Title: %s, Author: %s, Year: %d\n", book.BookId, book.Title, book.Author, book.YearPublished)
	}

	// Borrow a book
	borrowBookReq := &pb.BorrowBookRequest{
		BookId: addBookRes.BookId,
		UserId: "user123",
	}
	borrowBookRes, err := client.BorrowBook(context.Background(), borrowBookReq)
	if err != nil {
		log.Fatalf("BorrowBook failed: %v", err)
	}
	fmt.Printf("Borrow Success: %v\n", borrowBookRes.Success)
}
