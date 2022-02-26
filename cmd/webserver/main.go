package main

import (
	"log"
	"net/http"
	"webservice/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Router)
	log.Println("Listening...")
	err := http.ListenAndServe(":3000", mux)

	// http.HandleFunc("/", handlers.Router)

	// err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
