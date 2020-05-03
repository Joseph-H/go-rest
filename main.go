package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book ...
type Book struct {
	ID     string
	Isbn   string
	Title  string
	Author *Author
}

// Author ...
type Author struct {
	Firstname string
	Lastname  string
}

var books []Book

func getBooks(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

func getBook(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	params := mux.Vars(req) // get the params on the request

	for _, book := range books {
		if book.ID == params["id"] {
			encoder.Encode(book)
			return
		}
	}

	// if book is not found
	encoder.Encode(&Book{})
}

func createBook(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book Book

	json.NewDecoder(req.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(writer).Encode(book)
}

func updateBook(writer http.ResponseWriter, req *http.Request) {

}

func deleteBook(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	param := mux.Vars(req)

	for index, book := range books {
		if book.ID == param["id"] {
			// deleting a book from the middle
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(writer).Encode(books)
}

func main() {
	router := mux.NewRouter()

	// Mock data
	books = []Book{
		Book{
			ID:    "1",
			Isbn:  "1234",
			Title: "Radical Candor",
			Author: &Author{
				Firstname: "Kim",
				Lastname:  "Scott",
			},
		},
	}

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("UPDATE")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
