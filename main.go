package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// init env
	envErr := InitEnv()
	if envErr != nil {
		log.Fatal("Error loading .env file")
		return
	}
	
	// init db
	store, dbErr := NewDB()
	if dbErr != nil {
		fmt.Println(dbErr)
		return
	}

	server := ApiServer{address: ":5000", store: store}
	server.Run()
}



