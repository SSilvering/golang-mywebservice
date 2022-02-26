package main

import (
	"log"
	"net/http"
	"webservice/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Router)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
