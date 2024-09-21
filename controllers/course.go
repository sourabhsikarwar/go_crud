package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	mydatabase "github.com/sourabhsikarwar/go_crud/db"
	"github.com/sourabhsikarwar/go_crud/models"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my courses API</h1>"))
}

// Get all courses
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(mydatabase.MyCourses)
}

// Get course
func GetCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range mydatabase.MyCourses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
	return
}

// Create course
func CreateCourse(w http.ResponseWriter, r *http.Request) {
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

	mydatabase.MyCourses = append(mydatabase.MyCourses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// Update course
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, course := range mydatabase.MyCourses {
		if course.CourseID == params["id"] {
			mydatabase.MyCourses = append(mydatabase.MyCourses[:i], mydatabase.MyCourses[i+1:]...)

			var course models.Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			if course.IsEmpty() {
				json.NewEncoder(w).Encode("Course data is empty!")
				return
			}
			course.CourseID = params["id"]
			mydatabase.MyCourses = append(mydatabase.MyCourses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course ID not found")
	return
}

// Delete course
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, course := range mydatabase.MyCourses {
		if course.CourseID == params["id"] {
			mydatabase.MyCourses = append(mydatabase.MyCourses[:i], mydatabase.MyCourses[i+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("Course ID not found")
	return
}
