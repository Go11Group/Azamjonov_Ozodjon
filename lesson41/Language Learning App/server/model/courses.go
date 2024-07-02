package model

type Course struct {
	CourseId    string `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Time
}

type CourseFilter struct {
	Title       *string
	Description *string
	Limit       *string
	Offset      *string
}

type PopularCourse struct {
	CourseId    string `json:"course_id"`
	Title       string `json:"title"`
	Enrollments_count string `json:"enrollments_count"`
}