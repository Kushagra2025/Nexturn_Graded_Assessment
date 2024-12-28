package handlers

import (
	"encoding/json"
	"fmt"
	"http_server_exercise/students"
	"net/http"
	"sort"
	"strconv"
)

// Adding Students data
var Students_data = []students.Student{
	{ID: 1, Name: "John Doe", Subject: "CSE", GPA: 3, Major: "CSE"},
	{ID: 2, Name: "Jane Doe", Subject: "Mathematics", GPA: 4, Major: "CSE and AI"},
	{ID: 3, Name: "Erik Ten Hag", Subject: "Sports", GPA: 1, Major: "Sports and Tech"},
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Students_data)
}

// Default Route Handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the GO HTTP Server.")
}

// About Route Handler
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to About page")
}

// Contact Route Handler
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Contact page")
}

// GetStudentByID Route handler
func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	for _, student := range Students_data {
		if student.ID == id {
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	http.Error(w, "Student Not Found!", http.StatusNotFound)
}

// GetStudentByName Route handler
func GetStudentByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")

	var foundStudents []students.Student
	for _, student := range Students_data {
		if student.Name == name {
			foundStudents = append(foundStudents, student)
		}
	}

	if len(foundStudents) > 0 {
		json.NewEncoder(w).Encode(foundStudents)
	} else {
		http.Error(w, "No Student Found with the given Name", http.StatusNotFound)
	}
}

// GetStudentByMajor Route handler
func GetStudentByMajor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	major := r.URL.Query().Get("major")

	var foundStudents []students.Student
	for _, student := range Students_data {
		if student.Major == major {
			foundStudents = append(foundStudents, student)
		}
	}

	if len(foundStudents) > 0 {
		json.NewEncoder(w).Encode(foundStudents)
	} else {
		http.Error(w, "No Student Found with the given Major", http.StatusNotFound)
	}
}

// GetSortedStudents Route handler
func GetSortedStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sortOrder := r.URL.Query().Get("sort")

	// Sorting students based on GPA
	if sortOrder == "asc" {
		sort.Slice(Students_data, func(i, j int) bool {
			return Students_data[i].GPA < Students_data[j].GPA
		})
	} else if sortOrder == "desc" {
		sort.Slice(Students_data, func(i, j int) bool {
			return Students_data[i].GPA > Students_data[j].GPA
		})
	} else {
		http.Error(w, "Invalid sort parameter. Use 'asc' or 'desc'.", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(Students_data)
}

// AddStudent Route handler
func AddStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newStudent students.Student

	if err := json.NewDecoder(r.Body).Decode(&newStudent); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Generate a new ID for the student
	newStudent.ID = len(Students_data) + 1
	Students_data = append(Students_data, newStudent)

	json.NewEncoder(w).Encode(newStudent)
}

// DeleteStudent Route handler
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	for i, student := range Students_data {
		if student.ID == id {
			Students_data = append(Students_data[:i], Students_data[i+1:]...)
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	http.Error(w, "Student Not Found!", http.StatusNotFound)
}

// UpdateStudent Route handler
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	var updatedStudent students.Student
	if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	for i, student := range Students_data {
		if student.ID == id {
			Students_data[i] = updatedStudent
			json.NewEncoder(w).Encode(updatedStudent)
			return
		}
	}
	http.Error(w, "Student Not Found!", http.StatusNotFound)
}
