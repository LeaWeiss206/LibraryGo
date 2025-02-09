package handler

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
    "LibraryGo/internal/model"
    "LibraryGo/internal/router"
    "github.com/gorilla/mux"
)

func TestAddBook(t *testing.T) {
    tests := []struct {
        name       string
        book       model.Book
        wantStatus int
        wantSuccess bool
    }{
        {
            name: "Valid Book",
            book: model.Book{
                Title:         "Test Book",
                Author:        "Test Author",
                PublishedYear: 2024,
            },
            wantStatus: http.StatusCreated,
            wantSuccess: true,
        },
        {
            name: "Invalid Book - No Title",
            book: model.Book{
                Author:        "Test Author",
                PublishedYear: 2024,
            },
            wantStatus: http.StatusBadRequest,
            wantSuccess: false,
        },
        {
            name: "Invalid Book - No Author",
            book: model.Book{
                Title:         "Test Book",
                PublishedYear: 2024,
            },
            wantStatus: http.StatusBadRequest,
            wantSuccess: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            r := router.SetupRouter()
            body, _ := json.Marshal(tt.book)
            req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
            req.Header.Set("Content-Type", "application/json")
            w := httptest.NewRecorder()
            r.ServeHTTP(w, req)

            if w.Code != tt.wantStatus {
                t.Errorf("Expected status code %d but got %d", tt.wantStatus, w.Code)
            }

            var response model.APIResponse
            err := json.Unmarshal(w.Body.Bytes(), &response)
            if err != nil {
                t.Fatalf("Failed to decode response: %v", err)
            }

            if response.Success != tt.wantSuccess {
                t.Errorf("Expected success to be %v but got %v", tt.wantSuccess, response.Success)
            }

            if response.Status.Code != tt.wantStatus {
                t.Errorf("Expected status code %d but got %d", tt.wantStatus, response.Status.Code)
            }

            if tt.wantSuccess {
                book, ok := response.Data.(map[string]interface{})
                if !ok {
                    t.Fatal("Expected book data in response")
                }
                if book["title"] != tt.book.Title {
                    t.Errorf("Expected title %s but got %s", tt.book.Title, book["title"])
                }
            }
        })
    }
}

func TestGetBooks(t *testing.T) {
    // Setup test data
    r := router.SetupRouter()
    setupTestBooks(t, r)

    tests := []struct {
        name       string
        url        string
        wantStatus int
        wantCount  int
        wantSuccess bool
    }{
        {
            name:       "Get All Books",
            url:        "/books",
            wantStatus: http.StatusOK,
            wantCount:  3,
            wantSuccess: true,
        },
        {
            name:       "Filter By Author",
            url:        "/books?author=Test Author 1",
            wantStatus: http.StatusOK,
            wantCount:  1,
            wantSuccess: true,
        },
        {
            name:       "Filter By Year Range",
            url:        "/books?startYear=2023&endYear=2024",
            wantStatus: http.StatusOK,
            wantCount:  2,
            wantSuccess: true,
        },
        {
            name:       "Invalid Year Format",
            url:        "/books?startYear=invalid",
            wantStatus: http.StatusBadRequest,
            wantCount:  0,
            wantSuccess: false,
        },
        {
            name:       "No Results",
            url:        "/books?author=Nonexistent Author",
            wantStatus: http.StatusOK,
            wantCount:  0,
            wantSuccess: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req, _ := http.NewRequest("GET", tt.url, nil)
            w := httptest.NewRecorder()
            r.ServeHTTP(w, req)

            if w.Code != tt.wantStatus {
                t.Errorf("Expected status code %d but got %d", tt.wantStatus, w.Code)
            }

            var response model.APIResponse
            err := json.Unmarshal(w.Body.Bytes(), &response)
            if err != nil {
                t.Fatalf("Failed to decode response: %v", err)
            }

            if response.Success != tt.wantSuccess {
                t.Errorf("Expected success to be %v but got %v", tt.wantSuccess, response.Success)
            }

            if tt.wantSuccess && response.Meta != nil {
                if response.Meta.Count != tt.wantCount {
                    t.Errorf("Expected %d books but got %d", tt.wantCount, response.Meta.Count)
                }
            }
        })
    }
}

func TestGetBookByID(t *testing.T) {
    r := router.SetupRouter()
    setupTestBooks(t, r)

    tests := []struct {
        name       string
        bookID     string
        wantStatus int
        wantSuccess bool
    }{
        {
            name:       "Valid ID",
            bookID:    "1",
            wantStatus: http.StatusOK,
            wantSuccess: true,
        },
        {
            name:       "Invalid ID Format",
            bookID:    "invalid",
            wantStatus: http.StatusBadRequest,
            wantSuccess: false,
        },
        {
            name:       "Not Found",
            bookID:    "999",
            wantStatus: http.StatusNotFound,
            wantSuccess: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req, _ := http.NewRequest("GET", fmt.Sprintf("/books/%s", tt.bookID), nil)
            w := httptest.NewRecorder()
            r.ServeHTTP(w, req)

            if w.Code != tt.wantStatus {
                t.Errorf("Expected status code %d but got %d", tt.wantStatus, w.Code)
            }

            var response model.APIResponse
            err := json.Unmarshal(w.Body.Bytes(), &response)
            if err != nil {
                t.Fatalf("Failed to decode response: %v", err)
            }

            if response.Success != tt.wantSuccess {
                t.Errorf("Expected success to be %v but got %v", tt.wantSuccess, response.Success)
            }
        })
    }
}

func TestDeleteBookByID(t *testing.T) {
    r := router.SetupRouter()
    setupTestBooks(t, r)

    tests := []struct {
        name       string
        bookID     string
        wantStatus int
        wantSuccess bool
    }{
        {
            name:       "Valid ID",
            bookID:    "1",
            wantStatus: http.StatusNoContent,
            wantSuccess: true,
        },
        {
            name:       "Invalid ID Format",
            bookID:    "invalid",
            wantStatus: http.StatusBadRequest,
            wantSuccess: false,
        },
        {
            name:       "Not Found",
            bookID:    "999",
            wantStatus: http.StatusNotFound,
            wantSuccess: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req, _ := http.NewRequest("DELETE", fmt.Sprintf("/books/%s", tt.bookID), nil)
            w := httptest.NewRecorder()
            r.ServeHTTP(w, req)

            if w.Code != tt.wantStatus {
                t.Errorf("Expected status code %d but got %d", tt.wantStatus, w.Code)
            }

            if tt.wantStatus != http.StatusNoContent {
                var response model.APIResponse
                err := json.Unmarshal(w.Body.Bytes(), &response)
                if err != nil {
                    t.Fatalf("Failed to decode response: %v", err)
                }

                if response.Success != tt.wantSuccess {
                    t.Errorf("Expected success to be %v but got %v", tt.wantSuccess, response.Success)
                }
            }
        })
    }
}

// Helper function to setup test data
func setupTestBooks(t *testing.T, r *mux.Router) {
    testBooks := []model.Book{
        {
            Title:         "Test Book 1",
            Author:        "Test Author 1",
            PublishedYear: 2024,
        },
        {
            Title:         "Test Book 2",
            Author:        "Test Author 2",
            PublishedYear: 2023,
        },
        {
            Title:         "Test Book 3",
            Author:        "Test Author 3",
            PublishedYear: 2022,
        },
    }

    for _, book := range testBooks {
        body, _ := json.Marshal(book)
        req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
        req.Header.Set("Content-Type", "application/json")
        w := httptest.NewRecorder()
        r.ServeHTTP(w, req)

        if w.Code != http.StatusCreated {
            t.Fatalf("Failed to setup test data: %v", w.Body.String())
        }
    }
}