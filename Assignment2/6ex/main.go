package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User model represents a user in the database.
type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"`
	Age  int    `json:"age" gorm:"not null"`
}

var db *gorm.DB

// Setup the database connection
func setupDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres dbname=assignment2 password=040403 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err) // Log and exit on connection failure
	}
	return db
}

// Direct SQL Queries

// (Anna)
func getUsersSQL(c *gin.Context) {
	age := c.Query("age")
	query := "SELECT * FROM users"

	if age != "" {
		query += " WHERE age = $1"
	}

	// Execute the query
	rows, err := db.Raw(query, age).Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users) // Return the list of users
}

// Create User (POST /users)
func createUserSQL(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", user.Name, user.Age)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Name must be unique"})
		return
	}

	c.JSON(http.StatusCreated, user) // Return created user
}

// Update User (PUT /users/{id}) (Anna)
func updateUserSQL(c *gin.Context) {
	id := c.Param("id") // Get user ID from URL parameters
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Exec("UPDATE users SET name = $1, age = $2 WHERE id = $3", user.Name, user.Age, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user) // Return updated user
}

// Delete User (DELETE /users/{id}) (Anna)
func deleteUserSQL(c *gin.Context) {
	id := c.Param("id") // Get user ID from URL parameters
	err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.Status(http.StatusNoContent) // Return no content status
}

// GORM Routes

// Get Users (GET /users)
// This handler fetches all users using GORM.
func getUsersGORM(c *gin.Context) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // Handle error
		return
	}
	c.JSON(http.StatusOK, users) // Return the list of users
}

// Create User (POST /users)
// This handler inserts a new user into the database using GORM.
func createUserGORM(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Handle input validation error
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Name must be unique"}) // Handle unique constraint violation
		return
	}

	c.JSON(http.StatusCreated, user) // Return created user
}

// Update User (PUT /users/{id})
// This handler updates an existing user by ID using GORM.
func updateUserGORM(c *gin.Context) {
	id := c.Param("id") // Get user ID from URL parameters
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"}) // Handle not found error
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Handle input validation error
		return
	}

	db.Save(&user)              // Save updated user
	c.JSON(http.StatusOK, user) // Return updated user
}

// Delete User (DELETE /users/{id})
// This handler deletes a user by ID using GORM.
func deleteUserGORM(c *gin.Context) {
	id := c.Param("id") // Get user ID from URL parameters
	if err := db.Delete(&User{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"}) // Handle not found error
		return
	}

	c.Status(http.StatusNoContent) // Return no content status
}

// Main function
func main() {
	db = setupDatabase() // Set up the database connection

	// Set up Gin router
	router := gin.Default()

	// Direct SQL routes
	router.GET("/sql/users", getUsersSQL)          // Route to get users using direct SQL
	router.POST("/sql/users", createUserSQL)       // Route to create a user using direct SQL
	router.PUT("/sql/users/:id", updateUserSQL)    // Route to update a user using direct SQL
	router.DELETE("/sql/users/:id", deleteUserSQL) // Route to delete a user using direct SQL

	// GORM routes
	router.GET("/gorm/users", getUsersGORM)          // Route to get users using GORM
	router.POST("/gorm/users", createUserGORM)       // Route to create a user using GORM
	router.PUT("/gorm/users/:id", updateUserGORM)    // Route to update a user using GORM
	router.DELETE("/gorm/users/:id", deleteUserGORM) // Route to delete a user using GORM

	// Start the server
	router.Run(":8080") // Start the HTTP server on port 8080
}
