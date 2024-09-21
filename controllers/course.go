package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sourabhsikarwar/go_crud/models"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my courses API</h1>"))
}

// Get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(MyCourses)
}

// Get course
func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range MyCourses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
	return
}

// Create course
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a course")
	w.Header().Set("Content-Type", "application/json")

	// Failure cases
	if r.Body == nil {
		json.NewEncoder(w).Encode("No data found in request body!")
		return
	}

	var course models.Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Course data is empty!")
		return
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	course.CourseID = strconv.Itoa(rand.Intn(100))

	MyCourses = append(MyCourses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// Update course
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, course := range MyCourses {
		if course.CourseID == params["id"] {
			MyCourses = append(MyCourses[:i], MyCourses[i+1:]...)

			var course models.Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			if course.IsEmpty() {
				json.NewEncoder(w).Encode("Course data is empty!")
				return
			}
			course.CourseID = params["id"]
			MyCourses = append(MyCourses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course ID not found")
	return
}

// Delete course
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, course := range MyCourses {
		if course.CourseID == params["id"] {
			MyCourses = append(MyCourses[:i], MyCourses[i+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("Course ID not found")
	return
}
