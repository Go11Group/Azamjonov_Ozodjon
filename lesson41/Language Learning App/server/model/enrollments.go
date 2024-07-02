package model

import "time"

type Enrollment struct {
	EnrollmentId   string    `json:"enrollment_id"`
	UserId         string    `json:"user_id"`
	CourseId       string    `json:"course_id"`
	EnrollmentDate time.Time `json:"enrollment_date"`
	Time
}

type EnrollmentFilter struct {
	UserId         *string
	CourseId       *string
	EnrollmentDate *time.Time
	Limit          *string
	Offset         *string
}
