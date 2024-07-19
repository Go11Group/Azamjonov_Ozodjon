package main

import (
	"github.com/imtihon/handler"
	"github.com/imtihon/storage/postgres"
	"log"
)

func main() {
	// Connect to PostgreSQL database
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Ensure database connection is closed when main function exits

	// Initialize repositories using the connected database
	userRepo := postgres.NewUserRepo(db)
	courseRepo := postgres.NewCourseRepo(db)
	lessonRepo := postgres.NewLessonRepo(db)
	enrollmentRepo := postgres.NewEnrollmentRepo(db)

	// Initialize handler with repositories
	h := &handler.Handler{
		Users:       userRepo,
		Courses:     courseRepo,
		Lessons:     lessonRepo,
		Enrollments: enrollmentRepo,
	}

	// Start HTTP server with the handler
	server := handler.Handlerr(h)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
