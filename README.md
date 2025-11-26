
Here’s your full README in that format:

---

# 📚 Go + React Library Management System

Library Management System is a full‑stack capstone project built with **Go (backend)** and **React + TailwindCSS (frontend)**. It provides a dashboard interface for managing books, allowing users to add, update, delete, borrow, and return books. The backend exposes RESTful JSON APIs, while the frontend offers a responsive, styled interface.

---

## Features

- Add new books with title, author, and year  
- List all books in a dashboard table  
- Borrow and return books with status badges  
- Update book details (title, author, year)  
- Delete books from the system  
- Error handling and loading states  
- Responsive TailwindCSS design  
- Backend unit tests with `go test`

---

## Tech Stack

- **Go 1.20+** – Backend API  
- **React 18+** – Frontend UI  
- **TailwindCSS** – Styling framework  
- **Jest + React Testing Library** (optional) – Frontend testing  
- **SQLite/PostgreSQL** (future option) – Persistent storage  

---

## Setup Instructions

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd Go-react-Capstone

2. **Start Backend Server**
    ```bash
    cd Backend
    go run main.go

Server runs at: http://localhost:8080

3. **Start the Frontend (React + Vite)**
   ```bash
   cd ../frontend
   npm install
   npm run dev

## Api Endpoints
- api/books 
- api/add
- api/borrow?id={id}
- api/return?id={id}
- pi/update?id={id}
- api/delete?id={id}


## Project Structure
Go-react-Capstone/
├── Backend/
│   ├── main.go         # Go backend server
│   ├── main_test.go    # Unit tests for API endpoints
│   └── ...
├── frontend/
│   ├── src/App.jsx     # React dashboard
│   ├── src/main.jsx    # React entry point
│   ├── index.html      # Tailwind CDN or build entry
│   └── ...
└── README.md

## Running Tests
Backend Tests
Bashcd Backend
go test -v

**Example Output**
=== RUN   TestAddAndListBooks
--- PASS: TestAddAndListBooks (0.00s)
=== RUN   TestBorrowAndReturnBook
--- PASS: TestBorrowAndReturnBook (0.00s)
=== RUN   TestUpdateBook
--- PASS: TestUpdateBook (0.00s)
=== RUN   TestDeleteBook
--- PASS: TestDeleteBook (0.00s)
PASS
ok  	github.com/FrostyJames/Go-react-Capstone/backend	0.006s

## Frontend Tests (optional – add if you implement Jest)
Bashcd frontend
npm run test

## How to Use 
- Run Backend
cd Backend
go run main.go

- Run the Frontend
cd frontend
npm run dev
 

 **Navigate through the dashboard**
Library Management Dashboard Main View- Add new books
- Borrow / Return books
- Update book details
- Delete books
- View all records


## Notes & Limitations

Data is stored in-memory (resets when the Go server restarts)
Book updates currently use window.prompt() – ready to be replaced with a modal
No authentication (intentionally kept simple for capstone scope)


## Future Enhancements

Replace in-memory storage with SQLite or PostgreSQL
Add search and filter functionality
Implement proper modal forms for adding/updating books
Add user authentication and borrowing history
Write comprehensive frontend tests with Jest + RTL
Dockerize the full application for easy deployment


## Author
James Ivan
 Software Engineering Capstone Project
