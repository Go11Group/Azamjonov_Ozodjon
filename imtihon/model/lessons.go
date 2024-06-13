package model

import "time"

type Lessons struct {
	LessonId  string     `json:"lesson_id"`
	CourseId  string     `json:"course_id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
