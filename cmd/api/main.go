package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// ...
	// set application configuration
	var app application

	// read from commandline

	// connect to database

	app.Domain = "example.com"

	log.Printf("Starting server on port:%d", port)

	// start server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
