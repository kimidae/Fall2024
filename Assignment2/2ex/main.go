package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define a struct that represents the users table (Anna)
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Age  int    `gorm:"not null"`
}

// Connect to the PostgreSQL database
func connectDB() (*gorm.DB, error) {
	dsn := "host=localhost port=5432 user=postgres password=040403 dbname=assignment2 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// AutoMigrate the User model to create the users table (Anna)
func migrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

// Insert user into the database
func insertUser(db *gorm.DB, name string, age int) error {
	user := User{Name: name, Age: age}
	return db.Create(&user).Error
}

// Query and print all users from the database (Anna)
func queryUsers(db *gorm.DB) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}

func main() {
	// Connect to the database
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate to create the users table
	if err := migrateDB(db); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully")

	// Insert sample data
	users := []struct {
		name string
		age  int
	}{
		{"Anna", 30},
		{"Gregory", 25},
		{"John", 35},
	}

	for _, user := range users {
		if err := insertUser(db, user.name, user.age); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Users inserted successfully")

	// Query and print all users
	queryUsers(db)
}
