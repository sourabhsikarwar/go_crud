package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sourabhsikarwar/go_crud/controllers"
	mydatabase "github.com/sourabhsikarwar/go_crud/db"
)

func main() {
	fmt.Println("Hello, This is a go courses API!")

	// seeding data in MyCourses -- Fake DB
	mydatabase.SeedingData()

	r := mux.NewRouter()

	// Home API endpoint
	r.HandleFunc("/", controllers.ServeHome).Methods("GET")

	// Courses API endpoints
	r.HandleFunc("/api/courses", controllers.GetAllCourses).Methods("GET")
	r.HandleFunc("/api/courses/{id}", controllers.GetCourse).Methods("GET")
	r.HandleFunc("/api/courses", controllers.CreateCourse).Methods("POST")
	r.HandleFunc("/api/courses/{id}", controllers.UpdateCourse).Methods("PUT")
	r.HandleFunc("/api/courses/{id}", controllers.DeleteCourse).Methods("DELETE")

	// Start server and listen on port 8000
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
