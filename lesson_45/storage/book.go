package storage

import (
	"database/sql"
	"github.com/Azamjonov_Ozodjon/lesson_45/generator/library"
	"github.com/google/uuid"
)

type BookStorage struct {
	db *sql.DB
}

func NewBookStorage(db *sql.DB) *BookStorage {
	return &BookStorage{db: db}
}

func (s *BookStorage) AddBook(book *library.Book) (string, error) {
	bookID := uuid.New().String()
	query := `INSERT INTO books (book_id, title, author, year_published) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(query, bookID, book.Title, book.Author, book.YearPublished)
	if err != nil {
		return "", err
	}
	return bookID, nil
}

func (s *BookStorage) SearchBooks(query string) ([]*library.Book, error) {
	searchQuery := `SELECT book_id, title, author, year_published FROM books WHERE title ILIKE $1 OR author ILIKE $1`
	rows, err := s.db.Query(searchQuery, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*library.Book
	for rows.Next() {
		var book library.Book
		if err := rows.Scan(&book.BookId, &book.Title, &book.Author, &book.YearPublished); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func (s *BookStorage) BorrowBook(bookID string) (bool, error) {
	deleteQuery := `DELETE FROM books WHERE book_id = $1`
	result, err := s.db.Exec(deleteQuery, bookID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
