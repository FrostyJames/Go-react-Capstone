import { useEffect, useState } from "react";

function App() {
  const [books, setBooks] = useState([]);
  const [title, setTitle] = useState("");
  const [author, setAuthor] = useState("");
  const [year, setYear] = useState("");
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  // Fetch books
  useEffect(() => {
    fetch("http://localhost:8080/api/books")
      .then(res => {
        if (!res.ok) throw new Error("Failed to fetch books");
        return res.json();
      })
      .then(data => {
        setBooks(Array.isArray(data) ? data : []);
        setLoading(false);
      })
      .catch(err => {
        setError(err.message);
        setLoading(false);
      });
  }, []);

  // Add book
  const addBook = () => {
    if (!title || !author || !year) {
      setError("Please fill in all fields");
      return;
    }

    fetch("http://localhost:8080/api/add", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title, author, year: parseInt(year) })
    })
      .then(res => {
        if (!res.ok) throw new Error("Failed to add book");
        return res.json();
      })
      .then(newBook => {
        setBooks([...books, newBook]);
        setTitle(""); setAuthor(""); setYear("");
        setError(null);
      })
      .catch(err => setError(err.message));
  };

  // Borrow book
  const borrowBook = (id) => {
    fetch(`http://localhost:8080/api/borrow?id=${id}`)
      .then(res => {
        if (!res.ok) throw new Error("Failed to borrow book");
        return res.json();
      })
      .then(updated => setBooks(books.map(b => b.id === id ? updated : b)))
      .catch(err => setError(err.message));
  };

  // Return book
  const returnBook = (id) => {
    fetch(`http://localhost:8080/api/return?id=${id}`)
      .then(res => {
        if (!res.ok) throw new Error("Failed to return book");
        return res.json();
      })
      .then(updated => setBooks(books.map(b => b.id === id ? updated : b)))
      .catch(err => setError(err.message));
  };

  // Delete book
  const deleteBook = (id) => {
    fetch(`http://localhost:8080/api/delete?id=${id}`, { method: "DELETE" })
      .then(res => {
        if (!res.ok) throw new Error("Failed to delete book");
        return res.json();
      })
      .then(() => setBooks(books.filter(b => b.id !== id)))
      .catch(err => setError(err.message));
  };

  // Update book
  const updateBook = (id, current) => {
    const newTitle = prompt("Enter new title", current.title);
    const newAuthor = prompt("Enter new author", current.author);
    const newYear = prompt("Enter new year", current.year);

    if (!newTitle || !newAuthor || !newYear) {
      setError("Update cancelled or invalid input");
      return;
    }

    fetch(`http://localhost:8080/api/update?id=${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title: newTitle, author: newAuthor, year: parseInt(newYear) })
    })
      .then(res => {
        if (!res.ok) throw new Error("Failed to update book");
        return res.json();
      })
      .then(updatedBook => setBooks(books.map(b => b.id === id ? updatedBook : b)))
      .catch(err => setError(err.message));
  };

  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center p-6">
      <div className="w-full max-w-5xl bg-white shadow-lg rounded-lg p-6">
        <h1 className="text-3xl font-bold text-center mb-6">📚 Library Dashboard</h1>

        {/* Error message */}
        {error && <p className="text-red-500 mb-4">Error: {error}</p>}

        {/* Add book form */}
        <div className="flex gap-2 mb-6">
          <input
            className="flex-1 border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
            placeholder="Title"
            value={title}
            onChange={e => setTitle(e.target.value)}
          />
          <input
            className="flex-1 border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
            placeholder="Author"
            value={author}
            onChange={e => setAuthor(e.target.value)}
          />
          <input
            className="w-24 border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
            placeholder="Year"
            value={year}
            onChange={e => setYear(e.target.value)}
          />
          <button
            onClick={addBook}
            className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition"
          >
            Add
          </button>
        </div>

        {/* Loading state */}
        {loading ? (
          <p className="text-gray-500">Loading books...</p>
        ) : (
          <div className="overflow-x-auto">
            <table className="min-w-full border border-gray-200 rounded-lg">
              <thead className="bg-gray-200">
                <tr>
                  <th className="px-4 py-2 text-left">ID</th>
                  <th className="px-4 py-2 text-left">Title</th>
                  <th className="px-4 py-2 text-left">Author</th>
                  <th className="px-4 py-2 text-left">Year</th>
                  <th className="px-4 py-2 text-left">Status</th>
                  <th className="px-4 py-2 text-left">Actions</th>
                </tr>
              </thead>
              <tbody>
                {books.length === 0 ? (
                  <tr>
                    <td colSpan="6" className="text-center py-4 text-gray-500">
                      No books available. Add one above!
                    </td>
                  </tr>
                ) : (
                  books.map(book => (
                    <tr key={book.id} className="border-t">
                      <td className="px-4 py-2">{book.id}</td>
                      <td className="px-4 py-2 font-semibold">{book.title}</td>
                      <td className="px-4 py-2">{book.author}</td>
                      <td className="px-4 py-2">{book.year}</td>
                      <td className="px-4 py-2">
                        {book.borrowed ? (
                          <span className="bg-red-200 text-red-800 px-2 py-1 rounded">
                            Borrowed
                          </span>
                        ) : (
                          <span className="bg-green-200 text-green-800 px-2 py-1 rounded">
                            Available
                          </span>
                        )}
                      </td>
                      <td className="px-4 py-2 flex gap-2">
                        {book.borrowed ? (
                          <button
                            onClick={() => returnBook(book.id)}
                            className="bg-green-500 text-white px-3 py-1 rounded hover:bg-green-600 transition"
                          >
                            Return
                          </button>
                        ) : (
                          <button
                            onClick={() => borrowBook(book.id)}
                            className="bg-yellow-500 text-white px-3 py-1 rounded hover:bg-yellow-600 transition"
                          >
                            Borrow
                          </button>
                        )}
                        <button
                          onClick={() => updateBook(book.id, book)}
                          className="bg-blue-500 text-white px-3 py-1 rounded hover:bg-blue-600 transition"
                        >
                          Update
                        </button>
                        <button
                          onClick={() => deleteBook(book.id)}
                          className="bg-red-500 text-white px-3 py-1 rounded hover:bg-red-600 transition"
                        >
                          Delete
                        </button>
                      </td>
                    </tr>
                  ))
                )}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}

export default App;