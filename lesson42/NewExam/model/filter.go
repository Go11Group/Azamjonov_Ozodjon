package model

type Filter struct {
	Title          string
	Description    string
	UserId         string
	EnrollmentDate string
	CourseId       string
	Content        string
	Name           string
	Email          string
	Birthday       string
	Password       string
	Limit, Offset  int
}
