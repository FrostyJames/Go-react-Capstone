package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

// helper to decode JSON response
func decodeResponse(t *testing.T, rr *httptest.ResponseRecorder, v interface{}) {
    if err := json.NewDecoder(rr.Body).Decode(v); err != nil {
        t.Fatalf("Failed to decode response: %v", err)
    }
}

func TestAddAndListBooks(t *testing.T) {
    // reset state
    books = []Book{}
    nextID = 1

    // simulate adding a book
    body := []byte(`{"title":"1984","author":"George Orwell","year":1949}`)
    req := httptest.NewRequest("POST", "/api/add", bytes.NewBuffer(body))
    rr := httptest.NewRecorder()
    addBook(rr, req)

    if rr.Code != http.StatusOK {
        t.Fatalf("Expected 200, got %d", rr.Code)
    }

    var b Book
    decodeResponse(t, rr, &b)
    if b.Title != "1984" {
        t.Errorf("Expected title 1984, got %s", b.Title)
    }

    // simulate listing books
    req = httptest.NewRequest("GET", "/api/books", nil)
    rr = httptest.NewRecorder()
    listBooks(rr, req)

    var listed []Book
    decodeResponse(t, rr, &listed)
    if len(listed) != 1 {
        t.Errorf("Expected 1 book, got %d", len(listed))
    }
}

func TestBorrowAndReturnBook(t *testing.T) {
    books = []Book{{ID: 1, Title: "Test", Author: "A", Year: 2020}}
    nextID = 2

    // borrow
    req := httptest.NewRequest("GET", "/api/borrow?id=1", nil)
    rr := httptest.NewRecorder()
    borrowBook(rr, req)

    var b Book
    decodeResponse(t, rr, &b)
    if !b.Borrowed {
        t.Errorf("Expected borrowed=true, got false")
    }

    // return
    req = httptest.NewRequest("GET", "/api/return?id=1", nil)
    rr = httptest.NewRecorder()
    returnBook(rr, req)

    decodeResponse(t, rr, &b)
    if b.Borrowed {
        t.Errorf("Expected borrowed=false, got true")
    }
}

func TestUpdateBook(t *testing.T) {
    books = []Book{{ID: 1, Title: "Old", Author: "A", Year: 2000}}
    nextID = 2

    updated := Book{Title: "New", Author: "B", Year: 2021}
    body, _ := json.Marshal(updated)
    req := httptest.NewRequest("PUT", "/api/update?id=1", bytes.NewBuffer(body))
    rr := httptest.NewRecorder()
    updateBook(rr, req)

    var b Book
    decodeResponse(t, rr, &b)
    if b.Title != "New" {
        t.Errorf("Expected title New, got %s", b.Title)
    }
}

func TestDeleteBook(t *testing.T) {
    books = []Book{{ID: 1, Title: "DeleteMe", Author: "A", Year: 2000}}
    nextID = 2

    req := httptest.NewRequest("DELETE", "/api/delete?id=1", nil)
    rr := httptest.NewRecorder()
    deleteBook(rr, req)

    if len(books) != 0 {
        t.Errorf("Expected 0 books, got %d", len(books))
    }
}