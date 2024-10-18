package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Database configuration (Anna)
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "040403"
	dbname   = "assignment2"
)

// User struct for representing a user
type User struct {
	ID   int
	Name string
	Age  int
}

// Connect to the PostgreSQL database with connection pooling (Anna)
func connectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return sql.Open("postgres", psqlInfo)
}

// Create a users table with constraints (Anna)
func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) UNIQUE NOT NULL,
		age INT NOT NULL
	)`
	_, err := db.Exec(query)
	return err
}

// Insert users into the users table within a transaction (Anna)
func insertUsers(db *sql.DB, users []User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	for _, user := range users {
		_, err = tx.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", user.Name, user.Age)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// Query and print users with optional filtering and pagination (Anna)
func queryUsers(db *sql.DB, ageFilter int, page, pageSize int) ([]User, error) {
	var users []User
	query := "SELECT id, name, age FROM users WHERE age >= $1 LIMIT $2 OFFSET $3"
	rows, err := db.Query(query, ageFilter, pageSize, (page-1)*pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

// Update a user's details (Anna)
func updateUser(db *sql.DB, id int, name string, age int) error {
	_, err := db.Exec("UPDATE users SET name = $1, age = $2 WHERE id = $3", name, age, id)
	return err
}

// Delete a user by ID
func deleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func main() {
	// Connect to the database
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the users table
	if err := createTable(db); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Users table created successfully.")

	// Insert sample users
	users := []User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	if err := insertUsers(db, users); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Users inserted successfully.")

	// Query users with filtering and pagination
	ageFilter := 20
	page := 1
	pageSize := 2
	fetchedUsers, err := queryUsers(db, ageFilter, page, pageSize)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fetched Users:")
	for _, user := range fetchedUsers {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}

	// Update a user
	if err := updateUser(db, 1, "Alice Smith", 31); err != nil {
		log.Fatal(err)
	}
	fmt.Println("User updated successfully.")

	// Delete a user
	if err := deleteUser(db, 2); err != nil {
		log.Fatal(err)
	}
	fmt.Println("User deleted successfully.")
}
