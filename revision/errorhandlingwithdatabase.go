package revision

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	ID    int
	Name  string
	Email string
}

var db *sql.DB

func ConnectDB() (*sql.DB, error) {
	var err error
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/akash")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Connected to database..")
	return db, nil
}

// Createstudent (error handling)
func CreateStudent(db *sql.DB, name, email string) error {
	if db == nil {
		return fmt.Errorf("db is nil")
	}
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" || email == "" {
		return fmt.Errorf("name and email cannot be empty")
	}
	_, err := db.Exec("INSERT INTO students(name , email) VALUES(? , ?)", name, email)
	if err != nil {
		return fmt.Errorf("create failed: %v", err)
	}
	return nil
}

// insertstudent (error hadling)
func ReadStudent(db *sql.DB, id int) error {
	if db == nil {
		return fmt.Errorf("db is nil")
	}
	if id <= 0 {
		return fmt.Errorf("invalid student id")

	}
	var name, email string
	err := db.QueryRow("SELECT name , email FROM students WHERE id=?", id).Scan(&name, &email)
	if err == sql.ErrNoRows {
		return fmt.Errorf("student not found")
	}
	if err != nil {
		return fmt.Errorf("Read failed : %v", err)
	}
	fmt.Println("name:", name)
	fmt.Println("email:", email)
	return nil
}

// Updatestudent (error handling)
func UpdateStudent(db *sql.DB, id int, name, email string) error {
	if id <= 0 {
		return fmt.Errorf("invalid student id")
	}
	if name == "" || email == "" {
		return fmt.Errorf("name and email cannot be empty")
	}
	result, err := db.Exec("UPDATE students SET name=? , email=? WHERE id=?", name, email, id)
	if err != nil {
		return fmt.Errorf("Update failed: %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not check rows effected")
	}
	if rows == 0 {
		return fmt.Errorf("no student found to update")
	}
	return nil
}

// DeleteStudent (error handling)
func DeleteStudent(db *sql.DB, id int) error {
	if id <= 0 {
		return fmt.Errorf("Invalid student id")
	}
	result, err := db.Exec("DELETE FROM students WHERE id=?", id)
	if err != nil {
		return fmt.Errorf("Delete failed: %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not check rows affected")
	}
	if rows == 0 {
		return fmt.Errorf("student not found")
	}
	return nil
}

// main func
func ErrorHandlingWithDatabases() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("----CRUD operations----")
		fmt.Println("1.Create user")
		fmt.Println("2.Get user")
		fmt.Println("3.Update user")
		fmt.Println("4.Delete user")
		fmt.Println("5.Exit")
		fmt.Print("choose option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, email string
			fmt.Println("Enter name:")
			fmt.Scan(&name)
			fmt.Println("Enter email:")
			fmt.Scan(&email)

			if err := CreateStudent(db, name, email); err != nil {
				fmt.Println("error", err)
			} else {
				fmt.Println("student created succesfully!")
			}
			Pause()

		case 2:
			var id int
			fmt.Println("Enter Student id:")
			fmt.Scan(&id)

			if err := ReadStudent(db, id); err != nil {
				fmt.Println("Error:", err)
			}
			Pause()

		case 3:
			var id int
			var name, email string
			fmt.Println("Enter Student id:")
			fmt.Scan(&id)
			fmt.Println("Enter name:")
			fmt.Scan(&name)
			fmt.Println("Enter email:")
			fmt.Scan(&email)

			if err := UpdateStudent(db, id, name, email); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("student updated succesfully...")
			}
			Pause()

		case 4:
			var id int
			fmt.Println("Enter Student id:")
			fmt.Scan(&id)

			if err := DeleteStudent(db, id); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Student deleted succesfully...")
			}
			Pause()
		case 5:
			fmt.Println("Exiting succesfully...")
			return

		default:
			fmt.Println("Invalid choice!")
			Pause()
		}
		reader.ReadString('\n')
	}
}
