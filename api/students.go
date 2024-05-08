package handler

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	AcademicYear int    `json:"academic_year"`
}

var students []Student

func KestudentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newStudent Student
		json.NewDecoder(r.Body).Decode(&newStudent)

		// Check if the student ID already exists
		for _, student := range students {
			if student.ID == newStudent.ID {
				http.Error(w, "Student ID already exists", http.StatusBadRequest)
				return
			}
		}

		// If the student ID doesn't exist, add the student to the database
		students = append(students, newStudent)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newStudent)
	} else if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		name := r.URL.Query().Get("name")
		var foundStudents []Student

		if id != "" {
			// Search for students by name
			for _, student := range students {
				if student.ID == id {
					foundStudents = append(foundStudents, student)
				}
			}

			if len(foundStudents) > 0 {
				json.NewEncoder(w).Encode(foundStudents)
				return
			}
		}

		if name != "" {
			// Search for students by name
			for _, student := range students {
				if student.Name == name {
					foundStudents = append(foundStudents, student)
				}
			}

			if len(foundStudents) > 0 {
				json.NewEncoder(w).Encode(foundStudents)
				return
			}
		}
		// If no students found with the given name
		http.Error(w, "No students found", http.StatusNotFound)
	}
}

func Students() {
	// Initialize the router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/", KestudentsHandler).Methods("POST")
	router.HandleFunc("/", KestudentsHandler).Methods("GET")
	// Start the server
	http.ListenAndServe(":8080", router)
}
