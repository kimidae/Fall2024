package main

import (
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `gorm:"not null;unique"`
	Age     int     `gorm:"not null"`
	Profile Profile // One-to-one association
}

// Profile model
type Profile struct {
	ID                uint `gorm:"primaryKey"`
	UserID            uint `gorm:"not null;unique"`
	Bio               string
	ProfilePictureURL string
}

// Setup the database connection
func setupDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres dbname=assignment2 password=040403 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}
	return db
}

// AutoMigrate the models to create tables
func migrateModels(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		log.Fatal("failed to migrate models:", err)
	}
}

// Insert a User and Profile in a transaction
func insertUserWithProfile(db *gorm.DB, user User, profile Profile) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		profile.UserID = user.ID
		if err := tx.Create(&profile).Error; err != nil {
			return err
		}
		return nil
	})
}

// Query Users with Profiles using Eager Loading
func queryUsersWithProfiles(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Preload("Profile").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Update a User's Profile
func updateUserProfile(db *gorm.DB, userID uint, newProfile Profile) error {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	user.Profile = newProfile
	user.Profile.UserID = userID
	return db.Save(&user).Error
}

// Delete a User with Associated Profile
func deleteUserWithProfile(db *gorm.DB, userID uint) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", userID).Delete(&Profile{}).Error; err != nil {
			return err
		}
		return tx.Delete(&User{}, userID).Error
	})
}

func main() {
	db := setupDatabase()
	migrateModels(db)

	// Insert user and profile
	user := User{Name: "Alice", Age: 30}
	profile := Profile{Bio: "Software Developer", ProfilePictureURL: "https://www.example.com/picture.jpg"}
	if err := insertUserWithProfile(db, user, profile); err != nil {
		log.Println("Error inserting user and profile:", err)
	}

	// Query users with profiles
	users, err := queryUsersWithProfiles(db)
	if err != nil {
		log.Println("Error querying users:", err)
	} else {
		for _, u := range users {
			log.Printf("User: %s, Age: %d, Bio: %s\n", u.Name, u.Age, u.Profile.Bio)
		}
	}

	// Update profile
	newProfile := Profile{Bio: "Senior Developer", ProfilePictureURL: "https://www.example.com/newpicture.jpg"}
	if err := updateUserProfile(db, user.ID, newProfile); err != nil {
		log.Println("Error updating profile:", err)
	}

	// Delete user
	if err := deleteUserWithProfile(db, user.ID); err != nil {
		log.Println("Error deleting user:", err)
	}
}
