package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/harshRishi/go-bookstore/pkg/models"
	"github.com/harshRishi/go-bookstore/pkg/utils"
)

// GetBooks retrieves all books from the database and sends them in the response.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Failed to marshal books", http.StatusInternalServerError)
		log.Printf("Error marshaling books: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBookByID retrieves a book by its ID from the database and sends it in the response.
func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["bookId"]
	id, err := strconv.ParseInt(bookID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		log.Printf("Error parsing book ID: %v", err)
		return
	}

	book, db := models.GetBookById(id)
	if db.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		log.Printf("Error fetching book by ID: %v", db.Error)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to marshal book", http.StatusInternalServerError)
		log.Printf("Error marshaling book: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook creates a new book record in the database.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	if err := utils.ParseBody(r, newBook); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}

	book := newBook.CreateBook()
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to marshal book", http.StatusInternalServerError)
		log.Printf("Error marshaling book: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook deletes a book by its ID from the database.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["bookId"]
	id, err := strconv.ParseInt(bookID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		log.Printf("Error parsing book ID: %v", err)
		return
	}

	book := models.DeleteBook(id)
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to marshal book", http.StatusInternalServerError)
		log.Printf("Error marshaling book: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBook updates a book's details by its ID.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	if err := utils.ParseBody(r, updateBook); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}

	params := mux.Vars(r)
	bookID := params["bookId"]
	id, err := strconv.ParseInt(bookID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		log.Printf("Error parsing book ID: %v", err)
		return
	}

	bookDetails, db := models.GetBookById(id)
	if db.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		log.Printf("Error fetching book by ID: %v", db.Error)
		return
	}

	// Update fields if they are provided in the request
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publications != "" {
		bookDetails.Publications = updateBook.Publications
	}

	db.Save(bookDetails)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Failed to marshal book", http.StatusInternalServerError)
		log.Printf("Error marshaling book: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
