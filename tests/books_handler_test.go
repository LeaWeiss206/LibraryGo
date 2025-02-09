package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"LibraryGo/internal/model"
	"LibraryGo/internal/router"
)

func TestAddBook(t *testing.T) {
	r := router.SetupRouter()

	newBook := model.Book{
		Title:         "Test Book",
		Author:        "Test Author",
		PublishedYear: 2024,
	}

	body, _ := json.Marshal(newBook)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, w.Code)
	}

	var createdBook model.Book
	json.Unmarshal(w.Body.Bytes(), &createdBook)

	if createdBook.Title != newBook.Title || createdBook.Author != newBook.Author || createdBook.PublishedYear != newBook.PublishedYear {
		t.Errorf("Response body does not match request body")
	}
}

func TestGetAllBooks(t *testing.T) {
	r := router.SetupRouter()

	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}

func TestGetBookByID_NotFound(t *testing.T) {
	r := router.SetupRouter()

	req, _ := http.NewRequest("GET", "/books/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, w.Code)
	}
}

func TestDeleteBookByID_NotFound(t *testing.T) {
	r := router.SetupRouter()

	req, _ := http.NewRequest("DELETE", "/books/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, w.Code)
	}
}

