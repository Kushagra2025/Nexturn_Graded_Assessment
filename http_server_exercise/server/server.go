package server

import (
	"fmt"
	"http_server_exercise/handlers"
	"net/http"
)

func Start() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/students_data", handlers.GetAllStudents)

	http.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			//handlers.GetStudentByID(w, r)
			//handlers.GetStudentByMajor(w, r)
			// handlers.GetStudentByName(w, r)
			handlers.GetSortedStudents(w, r)
		case http.MethodPost:
			handlers.AddStudent(w, r)
		case http.MethodDelete:
			handlers.DeleteStudent(w, r)
		case http.MethodPut:
			handlers.UpdateStudent(w, r)
		default:
			http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		}
	})

	// Start on port 8080
	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server!!", err)
	}
}
