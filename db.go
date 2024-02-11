package main

import (
	"database/sql"
	"fmt"
	"go-todo/internal/database"
	"os"
)

type ApiConfig struct {
	DB *database.Queries
}

func NewDB() (ApiConfig, error) {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSslMode := os.Getenv("DB_SSL_MODE")

	connStr := fmt.Sprintf("user=%v dbname=%v password=%v sslmode=%v", dbUser, dbName, dbPassword, dbSslMode)
	conn, err := sql.Open("postgres", connStr)

	if err != nil {
		return ApiConfig{}, err
	}

	apiCfg := ApiConfig{
		DB: database.New(conn),
	}

	return apiCfg, nil
}