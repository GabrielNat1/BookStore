# 📚 Go API Project - Book Management  

This project is a REST API built with Go for managing a book catalog. It follows a modular structure to improve maintainability and scalability.  

## 📁 Project Structure  

```bash  
📂 cmd/  
│── main.go  -> Main entry point to start the application  
│── app.go   -> Application configuration  
📂 pkg/  
│── config/  -> Application settings (e.g., database)  
│── controllers/  
│   └── book_controller.go  -> Business logic for book-related endpoints  
│── models/  
│   └── book.go  -> Book data model definition  
│── routes/  
│   └── bookstore_routes.go  -> API route definitions  
│── utils/  
│   └── utils.go  -> Helper functions  
```  

## 🚀 Available Endpoints  

| Method  | Route             | Description              |  
|---------|------------------|--------------------------|  
| GET     | `/books`         | Retrieves all books     |  
| POST    | `/books`         | Creates a new book      |  
| GET     | `/books/{bookId}` | Retrieves a book by ID  |  
| PUT     | `/books/{bookId}` | Updates a book by ID    |  
| DELETE  | `/books/{bookId}` | Deletes a book by ID    |  

## 🛠 Technologies Used  

- Go (Golang)  
- Gin (for routing)  
- Database (e.g., PostgreSQL or MongoDB, depending on configuration)  

## 🏗 How to Run  

1. **Clone the repository**  
   ```bash
   git clone https://github.com/your-username/repo-name.git  
   cd repo-name  
   ```  
2. **Install dependencies**  
    ```bash  
   go mod tidy  
   ```  
3. **Run the application**  
    ```bash  
   go run cmd/main.go  
   ```  

## 📜 License  

This project is licensed under the [MIT](LICENSE) license.  

---
