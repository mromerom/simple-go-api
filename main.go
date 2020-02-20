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

	// informing client to expect JSON-formatted payload
	w.Header().Set("Content-Type", "application/json")

	// handle different request types
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))

	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))

	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))

	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	http.Handle("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
