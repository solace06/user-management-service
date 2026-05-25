package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {

	// Database configuration
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "khushi2001#"
	dbname := "users"

	// PostgreSQL connection string
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	// Create DB connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Verify DB connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to PostgreSQL")

	return db, nil
}