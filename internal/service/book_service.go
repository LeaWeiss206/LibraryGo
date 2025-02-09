package service

import (
	"LibraryGo/internal/model"
	"LibraryGo/internal/repository"
	"errors"
)

// BookService provides business logic
type BookService struct {
	repo *repository.BookRepository
}

// NewBookService initializes BookService
func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

// AddBook validates and adds a book
func (s *BookService) AddBook(book model.Book) (model.Book, error) {
	if book.Title == "" || book.Author == "" || book.PublishedYear <= 0 {
		return model.Book{}, errors.New("invalid book data")
	}

	return s.repo.AddBook(book), nil
}

// GetBookByID retrieves a book by ID
func (s *BookService) GetBookByID(id int) (model.Book, error) {
	return s.repo.GetBookByID(id)
}

// DeleteBookByID deletes a book
func (s *BookService) DeleteBookByID(id int) error {
	return s.repo.DeleteBookByID(id)
}

// GetBooksByAuthorAndYearRange retrieves books by author and published year range
func (s *BookService) GetBooks(author, startYear, endYear string) ([]model.Book, error) {
	return s.repo.GetBooks(author, startYear, endYear)
}


