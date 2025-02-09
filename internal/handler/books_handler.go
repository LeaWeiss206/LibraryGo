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
        utils.NewResponse().
            WithSuccess(false).
            WithError("INVALID_REQUEST", "Invalid request body", err.Error()).
            Send(w, http.StatusBadRequest)
        return
    }

    createdBook, err := h.service.AddBook(newBook)
    if err != nil {
        utils.NewResponse().
            WithSuccess(false).
            WithError("VALIDATION_ERROR", "Failed to create book", err.Error()).
            Send(w, http.StatusBadRequest)
        return
    }

    utils.NewResponse().
        WithSuccess(true).
        WithData(createdBook).
        Send(w, http.StatusCreated)
}

// GetBooks handles GET /books
func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
    author := r.URL.Query().Get("author")
    startYear := r.URL.Query().Get("startYear")
    endYear := r.URL.Query().Get("endYear")

    // Validate query parameters
    if startYear != "" && !utils.IsValidYear(startYear) {
        utils.NewResponse().
            WithSuccess(false).
            WithError("INVALID_PARAMETER", "Invalid startYear format", "Year must be a valid number").
            Send(w, http.StatusBadRequest)
        return
    }

    if endYear != "" && !utils.IsValidYear(endYear) {
        utils.NewResponse().
            WithSuccess(false).
            WithError("INVALID_PARAMETER", "Invalid endYear format", "Year must be a valid number").
            Send(w, http.StatusBadRequest)
        return
    }

    books, err := h.service.GetBooks(author, startYear, endYear)
    if err != nil {
        utils.NewResponse().
            WithSuccess(false).
            WithError("SERVER_ERROR", "Failed to retrieve books", err.Error()).
            Send(w, http.StatusInternalServerError)
        return
    }

    if len(books) == 0 {
        utils.NewResponse().
            WithSuccess(true).
            WithData([]model.Book{}).
            WithMeta(&model.MetaData{
                Total: 0,
                Count: 0,
            }).
            Send(w, http.StatusOK)
        return
    }

    utils.NewResponse().
        WithSuccess(true).
        WithData(books).
        WithMeta(&model.MetaData{
            Total: len(books),
            Count: len(books),
        }).
        Send(w, http.StatusOK)
}

// GetBookByID handles GET /books/{id}
func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    bookID, err := strconv.Atoi(params["id"])
    if err != nil {
        utils.NewResponse().
            WithSuccess(false).
            WithError("INVALID_ID", "Invalid book ID", "ID must be a valid number").
            Send(w, http.StatusBadRequest)
        return
    }

    book, err := h.service.GetBookByID(bookID)
    if err != nil {
        utils.NewResponse().
            WithSuccess(false).
            WithError("NOT_FOUND", "Book not found", "No book exists with the provided ID").
            Send(w, http.StatusNotFound)
        return
    }

    utils.NewResponse().
        WithSuccess(true).
        WithData(book).
        Send(w, http.StatusOK)
}

// DeleteBookByID handles DELETE /books/{id}
func (h *BookHandler) DeleteBookByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    bookID, err := strconv.Atoi(params["id"])
    if err != nil {
        utils.NewResponse().
            WithSuccess(false).
            WithError("INVALID_ID", "Invalid book ID", "ID must be a valid number").
            Send(w, http.StatusBadRequest)
        return
    }

    if err := h.service.DeleteBookByID(bookID); err != nil {
        utils.NewResponse().
            WithSuccess(false).
            WithError("NOT_FOUND", "Book not found", "No book exists with the provided ID").
            Send(w, http.StatusNotFound)
        return
    }

    utils.NewResponse().
        WithSuccess(true).
        Send(w, http.StatusNoContent)
}