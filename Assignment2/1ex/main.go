package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Setup PostgreSQL driver (Anna)
)

// Database connection parameters
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "040403"
	dbname   = "assignment2"
)

// Connects to the PostgreSQL database (Anna)
func connectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return sql.Open("postgres", connStr)
}

// Create the users table (Anna)
func createTable(db *sql.DB) error {
	query := `
	DROP TABLE users;
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	);`
	_, err := db.Exec(query)
	return err
}

// Insert data into the users table (Anna)
func insertUser(db *sql.DB, name string, age int) error {
	query := `INSERT INTO users (name, age) VALUES ($1, $2)`
	_, err := db.Exec(query, name, age)
	return err
}

// Query and print all users from the users table (Anna)
func queryUsers(db *sql.DB) error {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			return err
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	return rows.Err()
}

func main() {
	//(Anna)
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := createTable(db); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully")

	users := []struct {
		name string
		age  int
	}{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	//(Anna)
	for _, user := range users {
		if err := insertUser(db, user.name, user.age); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Users inserted successfully")

	if err := queryUsers(db); err != nil {
		log.Fatal(err)
	}
}
