package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)



func main() {
	connStr := "user=postgres dbname=go_todo password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
	}

	server := ApiServer{address: ":5000", store: Database{db}}
	server.Run()
}



