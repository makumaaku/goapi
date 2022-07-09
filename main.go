package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/makumaaku/goapi/model"
)

//構造体の作成

// Bookのデータを保持するスライスの作成
var books []model.Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Loop through books and find with id
	for _, item := range books {
			if item.ID == params["id"] {
					json.NewEncoder(w).Encode(item)
					return
			}
	}
	json.NewEncoder(w).Encode(&model.Book{})
}

// Create a Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000)) // Mock ID - not safe in production
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
			if item.ID == params["id"] {
					books = append(books[:index], books[index+1:]...)
					var book model.Book
					_ = json.NewDecoder(r.Body).Decode(&book)
					book.ID = params["id"]
					books = append(books, book)
					json.NewEncoder(w).Encode(book)
					return
			}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
			if item.ID == params["id"] {
					books = append(books[:index], books[index+1:]...)
					break
			}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
    // ルーターのイニシャライズ
    r := mux.NewRouter()

		books = append(books, model.Book{ID: "1", Title: "Book one",
		 Author: model.Author{FirstName: "Philip", LastName: "Williams"}})
    books = append(books, model.Book{ID: "2", Title: "Book Two", Author: model.Author{FirstName: "John", LastName: "Johnson"}})

    // ルート(エンドポイント)
    r.HandleFunc("/api/books", getBooks).Methods("GET")
    r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/api/books", createBook).Methods("POST")
    r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
    r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", r))
}