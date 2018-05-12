package engine

import (
	"net/http"
	"io"
	"log"
)

const PORT = 8080;

type EngineHandler struct {}

func (h *EngineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi " + r.URL.Path)
}

func NewEngine() {
	mux := http.NewServeMux()
	mux.Handle("/", &EngineHandler{})

	log.Printf("Server started on port %d", PORT)
	log.Fatal(http.ListenAndServe(":8080", mux))
}