package model

import (
	"time"
)

type Lesson struct {
	Id        string    `json:"id"`
	CourseId  string    `json:"course_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int       `json:"deleted_at"`
}
