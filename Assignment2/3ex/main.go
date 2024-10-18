package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User struct maps to the users table (Anna)
type User struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Age  int    `gorm:"not null" json:"age"`
}

// Connect to the PostgreSQL database
func connectDB() (*gorm.DB, error) {
	dsn := "host=localhost port=5432 user=postgres password=040403 dbname=assignment2 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// AutoMigrate the User model to create the users table
func migrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

// Get all users (Anna)
func getUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []User
		if err := db.Find(&users).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

// Create a new user (Anna)
func createUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := db.Create(&user).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}

// Update an existing user (Anna)
func updateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var user User
		if err := db.First(&user, id).Error; err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		db.Save(&user)
		json.NewEncoder(w).Encode(user)
	}
}

// Delete a user (Anna)
func deleteUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if err := db.Delete(&User{}, id).Error; err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	// Connect to the database (Anna)
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate to create the users table
	if err := migrateDB(db); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully")

	// Set up the router
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers(db)).Methods("GET")
	r.HandleFunc("/user", createUser(db)).Methods("POST")
	r.HandleFunc("/user/{id}", updateUser(db)).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUser(db)).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
