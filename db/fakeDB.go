package mydatabase

import "github.com/sourabhsikarwar/go_crud/models"

var MyCourses = []models.Course{}

func SeedingData() {
	MyCourses = append(
		MyCourses,
		models.Course{CourseID: "1", CourseName: "Go", CoursePrice: 100, Author: &models.Author{AuthorID: 1, AuthorName: "Sourabh"}},
		models.Course{CourseID: "2", CourseName: "Python", CoursePrice: 200, Author: &models.Author{AuthorID: 2, AuthorName: "Sourabh"}},
		models.Course{CourseID: "3", CourseName: "Java", CoursePrice: 300, Author: &models.Author{AuthorID: 3, AuthorName: "Sourabh"}},
		models.Course{CourseID: "4", CourseName: "C++", CoursePrice: 400, Author: &models.Author{AuthorID: 4, AuthorName: "Sourabh"}},
	)
}
