package models

import (
	"github.com/sourabhsikarwar/go_crud/models"
)

type Course struct {
	CourseID    int            `json:"course_id"`
	CourseName  string         `json:"course_name"`
	CoursePrice int            `json:"course_price"`
	Author      *models.Author `json:"author"`
}
