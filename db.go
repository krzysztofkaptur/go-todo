package main

import "database/sql"

type Database struct {
	db *sql.DB
}