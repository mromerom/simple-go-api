package main

import (
	// error logging
	"log"

	// to write REST API
	"net/http"
)

type server struct{}

// Handler interface
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// status 200; request successful
	w.WriteHeader(http.StatusOK)

	// informing client to expect JSON-formatted payload
	w.Header().Set("Content-Type", "application/json")

	// response
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
