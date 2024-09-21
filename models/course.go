package models

type Course struct {
	CourseID    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"course_price"`
	Author      *Author `json:"author"`
}

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}
