package main

import (
	"io"
	"net/http"
	"log"
	"regexp"
)

const PORT = 8000;

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

var validPath = regexp.MustCompile("^/(hello)/$")

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.URL.Path)
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

func main() {
	http.HandleFunc("/hello/", makeHandler(helloHandler))

	log.Printf("Server started on port %d", PORT)
	log.Fatal(http.ListenAndServe(":8080", nil))
}