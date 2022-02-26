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

	if err != nil {
		log.Fatal(err)
	}
}
