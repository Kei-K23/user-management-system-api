package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	dsn := os.Getenv("DB_CONNECTION_URL")
	if dsn == "" {
		log.Fatal("DB_CONNECTION_URL is not set")
	}

	// Parse the connection string
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Error when parsing db config: %v", err)
	}

	// Establish connection
	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Error when connecting to database: %v", err)
	}

	log.Println("Successfully connected to the database")
}
