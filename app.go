package main

import (
	"log"
	"net/http"
)

func main() {
	MigrateDatabase()
	
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

