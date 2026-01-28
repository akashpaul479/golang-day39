package revision

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Employee struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var DB *sql.DB

// connect database
func ConnectDataBase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/akash")
	if err != nil {
		fmt.Println("Error:", err)
	}
	if err := db.Ping(); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("database connection succesful...")
	return db, nil
}

// Create employee (error handling)
func CreateEmployee(db *sql.DB, name, email string) (int, error) {
	if db == nil {
		fmt.Println("db is nil")
	}
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" || email == "" {
		return 0, fmt.Errorf("name and email cannot be empty!")
	}

	result, err := db.Exec("INSERT INTO employees(name,email)VALUES(? , ?)", name, email)
	if err != nil {
		return 0, err
	}
	id, _ := result.LastInsertId()

	return int(id), nil
}

// Get employee (error handling)
func GetEmployee(db *sql.DB) ([]Employee, error) {
	if db == nil {
		fmt.Println("db is nil")
	}
	rows, err := db.Query("SELECT id, name, email FROM employees ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var e Employee
		rows.Scan(&e.Id, &e.Name, &e.Email)
		employees = append(employees, e)
	}
	return employees, nil
}

// Update employee (error handling)
func UpdateEmployee(db *sql.DB, id int, name, email string) error {
	if db == nil {
		return fmt.Errorf("db is nil")
	}
	if id <= 0 {
		return fmt.Errorf("invalid id")
	}
	result, err := db.Exec("UPDATE employees SET name=? , email=? WHERE id=?", name, email, id)
	if err != nil {
		return fmt.Errorf("update failed: %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not check rows efected")
	}
	if rows == 0 {
		return fmt.Errorf("employee not found ")
	}
	return nil
}

// Delete employee (error handling)
func DeleteEWmployee(db *sql.DB, id int) error {
	if db == nil {
		return fmt.Errorf("db id nil")
	}
	if id <= 0 {
		return fmt.Errorf("invalid id")
	}
	result, err := db.Exec("DELETE FROM employees WHERE id=?", id)
	if err != nil {
		return fmt.Errorf("Delete failed: %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not check rows affected")
	}
	if rows == 0 {
		return fmt.Errorf("employee not found")
	}
	return nil
}

// Create employee handler
func CreateEmployeeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e Employee
		json.NewDecoder(r.Body).Decode(&e)

		id, err := CreateEmployee(db, e.Name, e.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		e.Id = id
		json.NewEncoder(w).Encode(e)
	}
}

// Get employee handler
func GetEmployeeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employees, err := GetEmployee(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)
	}
}

// Update employee handler
func UpdateEmployeeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idstr := vars["id"]
		id, _ := strconv.Atoi(idstr)

		var e Employee
		json.NewDecoder(r.Body).Decode(&e)

		err := UpdateEmployee(db, id, e.Name, e.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		e.Id = id
		json.NewEncoder(w).Encode(e)
	}
}

// Delete employee handler
func DeleteEmployeeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idstr := vars["id"]
		id, _ := strconv.Atoi(idstr)

		err := DeleteEWmployee(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Employee deleted!"))
	}
}

// main func
func ErrorHandlingUsingServer() {
	db, err := ConnectDataBase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/employees", CreateEmployeeHandler(db)).Methods("POST")
	r.HandleFunc("/employees", GetEmployeeHandler(db)).Methods("GET")
	r.HandleFunc("/employees/{id}", UpdateEmployeeHandler(db)).Methods("PUT")
	r.HandleFunc("/employees/{id}", DeleteEmployeeHandler(db)).Methods("DELETE")

	fmt.Println("server running on port:8080")
	http.ListenAndServe(":8080", r)

}
