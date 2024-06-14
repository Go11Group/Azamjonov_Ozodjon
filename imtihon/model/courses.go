package model

import "time"

type Courses struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   int       `json:"deleted_at"`
}

type PopularCourse struct {
	CourseID         string `json:"course_id"`
	CourseTitle      string `json:"course_title"`
	EnrollmentsCount int    `json:"enrollments_count"`
}
