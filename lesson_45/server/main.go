package main

import (
	"context"
	"fmt"
	"net"
	"sync"

	pb "github.com/Azamjonov_Ozodjon/lesson_45/generator/library"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLibraryServiceServer
	mu    sync.Mutex
	books map[string]*pb.Book
}

func newServer() *server {
	return &server{
		books: make(map[string]*pb.Book),
	}
}

func (s *server) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	bookID := uuid.New().String()
	s.books[bookID] = &pb.Book{
		BookId:        bookID,
		Title:         req.Title,
		Author:        req.Author,
		YearPublished: req.YearPublished,
	}

	return &pb.AddBookResponse{BookId: bookID}, nil
}

func (s *server) SearchBook(ctx context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []*pb.Book
	for _, book := range s.books {
		if book.Title == req.Query || book.Author == req.Query {
			result = append(result, book)
		}
	}

	return &pb.SearchBookResponse{Books: result}, nil
}

func (s *server) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.books[req.BookId]; exists {
		delete(s.books, req.BookId)
		return &pb.BorrowBookResponse{Success: true}, nil
	}

	return &pb.BorrowBookResponse{Success: false}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := newServer()
	grpcServer := grpc.NewServer()
	pb.RegisterLibraryServiceServer(grpcServer, s)

	fmt.Println("Server is running on port :50051")
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
