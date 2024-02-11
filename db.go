package main

import (
	"database/sql"
	"fmt"
	"os"
)

type Database struct {
	db *sql.DB
}

func NewDB() (Database, error) {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSslMode := os.Getenv("DB_SSL_MODE")

	connStr := fmt.Sprintf("user=%v dbname=%v password=%v sslmode=%v", dbUser, dbName, dbPassword, dbSslMode)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return Database{}, err
	}

	return Database{db}, nil
}