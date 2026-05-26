package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {

	host := "localhost"
	port := 5432
	user := "postgres"
	password := "khushi2001#"
	dbname := "users"

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the db")

	return db, nil
}