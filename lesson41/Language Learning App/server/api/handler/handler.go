package handler

import (
	"database/sql"
	"language_learning_app/storage/postgres"
)

type Handler struct {
	Courses    *postgres.CourseRepo
	Enrollment *postgres.EnrollmentRepo
	Lessons    *postgres.LessonRepo
	Users      *postgres.UserRepo
}

func NewHandler(db *sql.DB) *Handler {
	courses := postgres.NewCourseRepo(db)
	enrollment := postgres.NewEnrollmentRepo(db)
	lessons := postgres.NewLessonRepo(db)
	users := postgres.NewUserRepo(db)

	return &Handler{
		Courses:    courses,
		Enrollment: enrollment,
		Lessons:    lessons,
		Users:      users,
	}
}
