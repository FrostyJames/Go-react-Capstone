Got it, James — let’s put everything we’ve built into a professional README.md. This will document your Go + React Capstone project, including setup, features, endpoints, frontend, and testing.
Here’s a complete draft you can drop into your repo:

📚 Go + React Library Management Dashboard
A full‑stack capstone project built with Go (backend) and React + TailwindCSS (frontend).
It implements a simple Library Management System with full CRUD operations, borrowing/returning functionality, and automated tests.

🚀 Features
- Backend (Go)
- Add, List, Borrow, Return, Update, Delete books
- In‑memory storage (resets on server restart)
- JSON API endpoints
- CORS enabled for frontend integration
- Unit tests with go test
- Frontend (React + TailwindCSS)
- Styled dashboard with table layout
- Add book form
- Borrow / Return toggle
- Update book (inline prompt)
- Delete book
- Error handling and loading states
- Responsive Tailwind design

🛠️ Backend Setup
Run the server
cd Backend
go run main.go


Server runs at:
http://localhost:8080


API Endpoints
|  |  |  | 
|  | /api/books |  | 
|  | /api/add |  | 
|  | /api/borrow?id={id} |  | 
|  | /api/return?id={id} |  | 
|  | /api/update?id={id} |  | 
|  | /api/delete?id={id} |  | 



🧪 Backend Testing
Tests are defined in main_test.go.
Run tests
cd Backend
go test ./...


Example verbose output
go test -v


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



🎨 Frontend Setup
Install dependencies
cd frontend
npm install


Run dev server
npm run dev


Frontend runs at:
http://localhost:5173


TailwindCSS
You can use either:
- CDN: Add <script src="https://cdn.tailwindcss.com"></script> in index.html
- Full install: npm install -D tailwindcss postcss autoprefixer && npx tailwindcss init -p

📊 Frontend Features
- Dashboard Table
- Displays all books with ID, Title, Author, Year, Status
- Status badges: green = Available, red = Borrowed
- Actions
- Borrow / Return toggle
- Update (via prompt)
- Delete (removes row)
- Add Book Form
- Inputs for Title, Author, Year
- Adds new book to backend and updates table

🔎 Testing Workflow
- Backend
- Run go test to validate API endpoints.
- Use curl or Postman to manually test endpoints.
- Frontend
- Run npm run dev and interact with the UI.
- Add, Borrow, Return, Update, Delete books.
- Confirm table updates correctly.

⚠️ Notes
- Books are stored in memory → data resets when backend restarts.
- For persistence, connect to a database (e.g., SQLite, PostgreSQL).
- Update currently uses prompt() — can be replaced with a Tailwind modal for better UX.

📌 Next Steps
- Add search & filter functionality in frontend.
- Replace prompt() with a modal form for updates.
- Connect backend to a database for persistence.
- Add frontend tests with Jest + React Testing Library.
- Deploy backend & frontend together (Docker or cloud).

👨‍💻 Author
James Ivan
Capstone project for Software Engineering — University of Eastern Africa, Baraton

👉 James, this README is ready to drop into your repo. Do you want me to also scaffold a frontend testing section (with Jest + React Testing Library examples) so your README covers both backend and frontend testing?
