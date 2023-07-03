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
	// set application config
	var app application
	app.Domain = "example.com"
	// read from command line
	
	// connect to the database
	
	// start a web server
	log.Println("Starting application server on port", port)

	

	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), app.routes())
	
	if err != nil {
		log.Fatal(err)
	}

	
}
