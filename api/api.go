package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
)

type Employee struct {
    ID       	 int    `json:"id"`
    Name     	 string `json:"name"`
    Lastname 	 string `json:"lastname"`
    Username 	 string `json:"username"`
    Email    	 string `json:"email"`
    Position 	 string `json:"position"`
	Address 	 string `json:"address"`
    DateOfBirth  string `json:"date_of_birth"`
    PlaceOfBirth string `json:"place_of_birth"`
}

func dbConnect() (*sql.DB, error) {
    dbUser := "srv_1"
    dbPassword := "srv_1"
    dbName := "admin_platform"
    dbHost := "localhost"
    dbPort := "3306"

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
    return sql.Open("mysql", dsn)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
    db, err := dbConnect()
    if err != nil {
        http.Error(w, "Error connecting to database: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    rows, err := db.Query("SELECT Id, Name, Lastname, Username, Email, Position, Address, DateOfBirth, PlaceOfBirth FROM employees")
    if err != nil {
        http.Error(w, "Error fetching employees: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var employees []Employee
    for rows.Next() {
        var employee Employee
        if err := rows.Scan(&employee.ID, &employee.Name, &employee.Lastname, &employee.Username, &employee.Email, &employee.Position, &employee.Address, &employee.DateOfBirth, &employee.PlaceOfBirth); err != nil {
            http.Error(w, "Error scanning employee: "+err.Error(), http.StatusInternalServerError)
            return
        }
        employees = append(employees, employee)
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(employees); err != nil {
        http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
    }
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
    db, err := dbConnect()
    if err != nil {
        http.Error(w, "Error connecting to database: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    var employee Employee
    if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
        http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
        return
    }

    _, err = db.Exec("INSERT INTO employees (Name, Lastname, Username, Email, Position, Address, DateOfBirth, PlaceOfBirth) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
        employee.Name, employee.Lastname, employee.Username, employee.Email, employee.Position, employee.Address, employee.DateOfBirth, employee.PlaceOfBirth)
    if err != nil {
        http.Error(w, "Error creating employee: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
    db, err := dbConnect()
    if err != nil {
        http.Error(w, "Error connecting to database: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    var employee Employee
    if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
        http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
        return
    }

    _, err = db.Exec("UPDATE employees SET Name=?, Lastname=?, Username=?, Email=?, Position=?, Address=?, DateOfBirth=?, PlaceOfBirth=? WHERE Id=?",
        employee.Name, employee.Lastname, employee.Username, employee.Email, employee.Position, employee.Address, employee.DateOfBirth, employee.PlaceOfBirth, employee.ID)
    if err != nil {
        http.Error(w, "Error updating employee: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
    db, err := dbConnect()
    if err != nil {
        http.Error(w, "Error connecting to database: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    var employee Employee
    if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
        http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
        return
    }

    _, err = db.Exec("DELETE FROM employees WHERE Id=?", employee.ID)
    if err != nil {
        http.Error(w, "Error deleting employee: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/employees", getEmployees)
    mux.HandleFunc("/api/employees/create", createEmployee)
    mux.HandleFunc("/api/employees/update", updateEmployee)
    mux.HandleFunc("/api/employees/delete", deleteEmployee)

    // Enable CORS
    corsOptions := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}), // Vue dev server
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type"}),
    )

    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", corsOptions(mux)))
}