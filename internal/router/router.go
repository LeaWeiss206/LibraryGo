package router

import (
	"LibraryGo/internal/handler"
	"LibraryGo/internal/repository"
	"LibraryGo/internal/service"

	"github.com/gorilla/mux"
)

// SetupRouter initializes the router
func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	repo := repository.NewBookRepository()
	bookService := service.NewBookService(repo)
	bookHandler := handler.NewBookHandler(bookService)

	r.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", bookHandler.GetBookByID).Methods("GET")
	r.HandleFunc("/books", bookHandler.AddBook).Methods("POST")
	r.HandleFunc("/books/{id}", bookHandler.DeleteBookByID).Methods("DELETE")

	return r
}