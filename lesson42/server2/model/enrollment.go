package model

import (
	"time"
)

type Enrollment struct {
	Id             string    `json:"id"`
	UserId         string    `json:"user_id"`
	CourseId       string    `json:"course_id"`
	EnrollmentDate string    `json:"enrollment_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      int       `json:"deleted_at"`
}
