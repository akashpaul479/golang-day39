package revision

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type User struct {
	Id    int
	Name  string
	Email string
}

var users = make(map[int]User)
var nextID = 1

// create error handling
func CreateUser(name, email string) error {
	if name == "" || email == "" {
		return errors.New("name and email cannot be empty")
	}
	for _, u := range users {
		if u.Email == email {
			return errors.New("email already exsists")
		}
	}
	user := User{
		Id:    nextID,
		Name:  name,
		Email: email,
	}
	users[nextID] = user
	nextID++

	return nil
}

// read error handling
func ReadUser(id int) (User, error) {
	if id <= 0 {
		return User{}, errors.New("invalid user id")
	}
	user, exsists := users[id]
	if !exsists {
		return User{}, errors.New("user not found ")
	}
	return user, nil
}

// update error handling
func Updateuser(id int, name string) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	if name == "" {
		return errors.New("name cannot be empty")
	}
	user, exsists := users[id]
	if !exsists {
		return errors.New("user not found")
	}
	user.Name = name
	users[id] = user
	return nil
}

// Delete error handling
func DeleteUser(id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	if _, exists := users[id]; !exists {
		return errors.New("user not found")
	}
	delete(users, id)

	return nil
}

// pause func
func Pause() {
	fmt.Println("Press enter to continue...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
func ErrorHandling() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("-----CRUD operations-----")
		fmt.Println("1.Create user")
		fmt.Println("2.Get user")
		fmt.Println("3.Update user")
		fmt.Println("4. Delete user")
		fmt.Println("5. Exit")
		fmt.Print("Choose option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, email string
			fmt.Println("Enter name:")
			fmt.Scan(&name)
			fmt.Println("Enter email:")
			fmt.Scan(&email)

			if err := CreateUser(name, email); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("user created succesfully")
			}
			Pause()
		case 2:
			var id int
			fmt.Println("Enter id:")
			fmt.Scan(&id)

			user, err := ReadUser(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("ID: %d, name: %s, Email:%s\n", user.Id, user.Name, user.Email)
			}
			Pause()
		case 3:
			var id int
			var name string
			fmt.Println("Enter user id:")
			fmt.Scan(&id)
			fmt.Println("Enter new name:")
			fmt.Scan(&name)

			if err := Updateuser(id, name); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Updated succesfully")
			}
			Pause()
		case 4:
			var id int
			fmt.Println("Enter id:")
			fmt.Scan(&id)
			if err := DeleteUser(id); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("deleted succesfully")
			}
			Pause()
		case 5:
			fmt.Println("Exiting succesfully...")

			return

		default:
			fmt.Println("Invalid choice")
			Pause()
		}
		reader.ReadString('\n')
	}
}
