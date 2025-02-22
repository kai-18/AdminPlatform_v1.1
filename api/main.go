package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func dbConnect() (*sql.DB, error) {
    dbUser := "root"
    dbPassword := "localhost"
    dbName := "test"
    dbHost := "localhost"
    dbPort := "3306"

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
    return sql.Open("mysql", dsn)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/employees", getEmployees)
    mux.HandleFunc("/api/employees/create", createEmployee)
    mux.HandleFunc("/api/employees/update", updateEmployee)
    mux.HandleFunc("/api/employees/delete", deleteEmployee)
	
    mux.HandleFunc("/api/users", getUsers)
    mux.HandleFunc("/api/users/create", createUser)
    mux.HandleFunc("/api/users/update", updateUser)
    mux.HandleFunc("/api/users/delete", deleteUser)

    // Enable CORS
    corsOptions := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type"}),
    )

    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", corsOptions(mux)))
}