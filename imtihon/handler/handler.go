package handler

import "github.com/imtihon/storage/postgres"

type Handler struct {
	Users       *postgres.UserRepo
	Courses     *postgres.CourseRepo
	Lessons     *postgres.LessonRepo
	Enrollments *postgres.EnrollmentRepo
}
