package repository

import (
	"LibraryGo/internal/model"
	"errors"
	"sync"
	"strconv"
)

// BookRepository manages book storage
type BookRepository struct {
	books  map[int]model.Book
	nextID int
	mu     sync.Mutex
}

// NewBookRepository initializes a book repository
func NewBookRepository() *BookRepository {
	return &BookRepository{
		books:  make(map[int]model.Book),
		nextID: 1,
	}
}

// AddBook saves a new book
func (repo *BookRepository) AddBook(book model.Book) model.Book {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	book.ID = repo.nextID
	repo.books[repo.nextID] = book
	repo.nextID++

	return book
}

// GetAllBooks retrieves all books
func (repo *BookRepository) GetAllBooks() []model.Book {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var bookList []model.Book
	for _, book := range repo.books {
		bookList = append(bookList, book)
	}
	return bookList
}

// GetBookByID retrieves a book by its ID
func (repo *BookRepository) GetBookByID(id int) (model.Book, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	book, exists := repo.books[id]
	if !exists {
		return model.Book{}, errors.New("book not found")
	}

	return book, nil
}

// DeleteBookByID removes a book
func (repo *BookRepository) DeleteBookByID(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.books[id]; !exists {
		return errors.New("book not found")
	}

	delete(repo.books, id)
	return nil
}

// GetBooks retrieves books by author and/or published year range
func (repo *BookRepository) GetBooks(author, startYear, endYear string) ([]model.Book, error) {
    repo.mu.Lock()
    defer repo.mu.Unlock()

    var filteredBooks []model.Book

    for _, book := range repo.books {
        // Filter by author if provided
        if author != "" && book.Author != author {
            continue
        }

        // Filter by startYear if provided
        if startYear != "" {
            startYearInt, err := strconv.Atoi(startYear)
            if err != nil {
                return nil, errors.New("invalid startYear format")
            }
            if book.PublishedYear < startYearInt {
                continue
            }
        }

        // Filter by endYear if provided
        if endYear != "" {
            endYearInt, err := strconv.Atoi(endYear)
            if err != nil {
                return nil, errors.New("invalid endYear format")
            }
            if book.PublishedYear > endYearInt {
                continue
            }
        }

        // If it passed all filters, add it to the result
        filteredBooks = append(filteredBooks, book)
    }

    // אם לא נמצאו ספרים, נחזור עם מערך ריק
    if len(filteredBooks) == 0 {
        return []model.Book{}, nil
    }

    return filteredBooks, nil
}
