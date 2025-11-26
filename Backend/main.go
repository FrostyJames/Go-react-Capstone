package main

import (
    "encoding/json"
    "net/http"
    "strconv"
)

type Book struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    Author   string `json:"author"`
    Year     int    `json:"year"`
    Borrowed bool   `json:"borrowed"`
}

var books []Book
var nextID = 1

// Utility: set common CORS headers
func setCORS(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// Add a book
func addBook(w http.ResponseWriter, r *http.Request) {
    setCORS(w)
    if r.Method == http.MethodOptions {
        return
    }

    var b Book
    if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    b.ID = nextID
    nextID++
    books = append(books, b)
    json.NewEncoder(w).Encode(b)
}

// List all books
func listBooks(w http.ResponseWriter, r *http.Request) {
    setCORS(w)
    if r.Method == http.MethodOptions {
        return
    }
    json.NewEncoder(w).Encode(books)
}

// Borrow a book
func borrowBook(w http.ResponseWriter, r *http.Request) {
    setCORS(w)
    if r.Method == http.MethodOptions {
        return
    }

    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)
    for i := range books {
        if books[i].ID == id {
            books[i].Borrowed = true
            json.NewEncoder(w).Encode(books[i])
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

// Return a book
func returnBook(w http.ResponseWriter, r *http.Request) {
    setCORS(w)
    if r.Method == http.MethodOptions {
        return
    }

    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)
    for i := range books {
        if books[i].ID == id {
            books[i].Borrowed = false
            json.NewEncoder(w).Encode(books[i])
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
    setCORS(w)
    if r.Method == http.MethodOptions {
        return
    }

    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)

    var updated Book
    if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    for i := range books {
        if books[i].ID == id {
            books[i].Title = updated.Title
            books[i].Author = updated.Author
            books[i].Year = updated.Year
            json.NewEncoder(w).Encode(books[i])
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
    setCORS(w)
    if r.Method == http.MethodOptions {
        return
    }

    idStr := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idStr)

    for i := range books {
        if books[i].ID == id {
            books = append(books[:i], books[i+1:]...)
            json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted"})
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

func main() {
    http.HandleFunc("/api/books", listBooks)
    http.HandleFunc("/api/add", addBook)
    http.HandleFunc("/api/borrow", borrowBook)
    http.HandleFunc("/api/return", returnBook)
    http.HandleFunc("/api/update", updateBook)   // ✅ new
    http.HandleFunc("/api/delete", deleteBook)   // ✅ new
    http.ListenAndServe(":8080", nil)
}