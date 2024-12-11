package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

// ConnectDB create a connection to the PostgreSQL database
func ConnectDB() {
	databaseUrl := os.Getenv("DATABASE_URL")
	var err error

	// Create connection pool
	Pool, err = pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Could not connect do PostgreSQL database: %v", err)
	} else {
		log.Println("Database connected")
	}

}
