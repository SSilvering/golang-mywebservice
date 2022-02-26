package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"time"
)

var regexExpresation, _ = regexp.Compile("[*]")

func Router(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-type") == "text/html" {
		fmt.Println(r.URL.Path)

		switch {
		case regexExpresation.MatchString(r.URL.Path):
			defaultHandler(w, r)
		default:
			defaultHandler(w, r)
		}
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request handled by default handler")
	content, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println("Request body cannot be read")
	}

	log.Println(string(content))

	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}
