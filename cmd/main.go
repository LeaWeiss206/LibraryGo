package main

import (
	"fmt"
	"log"
	"net/http"

	"LibraryGo/internal/router"
)

func main() {
	r := router.SetupRouter()

	port := ":8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}
