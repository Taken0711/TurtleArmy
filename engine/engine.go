package engine

import (
	"net/http"
	"io"
	"log"
)

const PORT = 8000;

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi " + r.URL.Path)
}

func NewEngine() {
	http.HandleFunc("/", helloHandler)

	log.Printf("Server started on port %d", PORT)
	log.Fatal(http.ListenAndServe(":8080", nil))
}