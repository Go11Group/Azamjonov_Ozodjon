package service

import (
	"context"
	"github.com/Azamjonov_Ozodjon/lesson_45/generator/library"
	"github.com/Azamjonov_Ozodjon/lesson_45/storage"
)

type server struct {
	library.UnimplementedLibraryServiceServer
	bookStorage *storage.BookStorage
}

func NewServer(bookStorage *storage.BookStorage) *server {
	return &server{
		bookStorage: bookStorage,
	}
}

func (s *server) AddBook(ctx context.Context, req *library.AddBookRequest) (*library.AddBookResponse, error) {
	book := &library.Book{
		Title:         req.Title,
		Author:        req.Author,
		YearPublished: req.YearPublished,
	}

	bookID, err := s.bookStorage.AddBook(book)
	if err != nil {
		return nil, err
	}

	return &library.AddBookResponse{BookId: bookID}, nil
}

func (s *server) SearchBook(ctx context.Context, req *library.SearchBookRequest) (*library.SearchBookResponse, error) {
	books, err := s.bookStorage.SearchBooks(req.Query)
	if err != nil {
		return nil, err
	}

	return &library.SearchBookResponse{Books: books}, nil
}

func (s *server) BorrowBook(ctx context.Context, req *library.BorrowBookRequest) (*library.BorrowBookResponse, error) {
	success, err := s.bookStorage.BorrowBook(req.BookId)
	if err != nil {
		return nil, err
	}

	return &library.BorrowBookResponse{Success: success}, nil
}
