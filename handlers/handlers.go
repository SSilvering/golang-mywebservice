package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

var defaultString := os.Getenv("DEFSTR")
var regexExpresation, _ := regexp.Compile("[*]")

func Router(w http.ResponseWriter, r *http.Request) {
	log.Println("URI: " + r.URL.Path)

	switch {
	case regexExpresation.MatchString(r.URL.Path):
		defaultHandler(w, r)
	default:
		defaultHandler(w, r)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-type") == "text/html" {
		log.Println("Request handled by default handler")
		content, err := io.ReadAll(r.Body)

		if err != nil {
			log.Println("Request body cannot be read")
		}

		log.Println("Request body: " + string(content))

		
		tm := time.Now().Format(time.RFC1123)

		if defaultString == "" {
		w.Write([]byte("The time is: " + tm))
		} else {
			w.Write([]byte(defaultString))
		}

	} else {
		log.Println("Received a request with a different content type. Type: " + r.Header.Get("Content-type"))
	}
}
