package helpers

import "github.com/sourabhsikarwar/go_crud/models"

func (c *models.Course) IsEmpty() {
	return c.CourseName == ""
}
