package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"LibraryGo/internal/model"
	"LibraryGo/internal/service"
	"LibraryGo/internal/utils"
)

// BookHandler handles HTTP requests
type BookHandler struct {
	service *service.BookService
}

// NewBookHandler creates a handler
func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// AddBook handles POST /books
func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	var newBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdBook, err := h.service.AddBook(newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}

// GetBooks handles GET /books
func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
    author := r.URL.Query().Get("author")
    startYear := r.URL.Query().Get("startYear")
    endYear := r.URL.Query().Get("endYear")

	if startYear != "" && !utils.IsValidYear(startYear) {
        http.Error(w, "Invalid startYear format", http.StatusBadRequest)
        return
    }

    if endYear != "" && !utils.IsValidYear(endYear) {
        http.Error(w, "Invalid endYear format", http.StatusBadRequest)
        return
    }

    books, err := h.service.GetBooks(author, startYear, endYear)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if len(books) == 0 {
        http.Error(w, "No books found", http.StatusNotFound) 
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}



// GetBookByID handles GET /books/{id}
func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := h.service.GetBookByID(bookID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// DeleteBookByID handles DELETE /books/{id}
func (h *BookHandler) DeleteBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteBookByID(bookID); err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}