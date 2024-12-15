package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB contains database connection instance
var DB *gorm.DB

// ConnectDB create a connection to the PostgreSQL database
func ConnectDB() {
	databaseUrl := os.Getenv("DATABASE_URL")
	var err error

	// Create connection pool
	DB, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect do PostgreSQL database: %v", err)
	} else {
		log.Println("Database connected")
	}

}
