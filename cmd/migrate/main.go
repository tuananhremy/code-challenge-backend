package main

import (
	"log"

	"code-challenge-backend/app"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize database
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&app.User{}, &app.Seat{}, &app.Booking{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
