package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Nguyen Phu Vinh Toan Golang RESTful API</h1>")
}

func Main() {
	// Initialize the router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/", Handler).Methods("GET")

	// Start the server
	http.ListenAndServe(":8080", router)
}
