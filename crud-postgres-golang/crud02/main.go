package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "teste"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	// Connect to the database
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to the database!")

	// Create a new user
	newUser := User{Name: "John Doe", Email: "john.doe@example.com"}
	err = createUser(db, &newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created user with ID", newUser.ID)

	// Read a user
	user, err := readUser(db, newUser.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read user:", user)

	// Update a user
	user.Name = "Jane Doe"
	err = updateUser(db, &user)
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated user")

	// Delete a user
	err = deleteUser(db, user.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted user")
}

// createUser creates a new user in the database and updates the user's ID
func createUser(db *sql.DB, user *User) error {
	err := db.QueryRow("INSERT INTO users(name, email) VALUES($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

// readUser reads a user from the database by ID
func readUser(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// updateUser updates a user in the database
func updateUser(db *sql.DB, user *User) error {
	_, err := db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// deleteUser deletes a user from the database
func deleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
