package model

// Book represents a book entity
type Book struct {
    ID            int `json:"id"`
    Title         string `json:"title"`
    Author        string `json:"author"`
    PublishedYear int    `json:"publishedYear"`
}