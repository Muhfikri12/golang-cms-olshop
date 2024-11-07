package main

import (
	"fmt"
	"net/http"

	"github.com/Muhfikri12/golang-cms-olshop/app/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/create-book", handler.FormCreateBook)
	mux.HandleFunc("/books", handler.CreateBook)
	mux.HandleFunc("/books-list", handler.ItemsList)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mux)
}
