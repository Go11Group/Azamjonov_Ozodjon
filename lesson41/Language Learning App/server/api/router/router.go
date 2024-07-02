package router

import (
	"language_learning_app/api/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServer(handler *handler.Handler) *http.Server {
	r := gin.Default()
	mainRouter := r.Group("/exam")

	// Create subRouter for Users
	userRouter := mainRouter.Group("/users")

	userRouter.POST("/create", handler.CreateUser)
	userRouter.GET("/getall", handler.GetUsers)
	userRouter.GET("/:id", handler.GetUserById)
	userRouter.GET("/:id/courses", handler.GetUsersCourses)
	userRouter.PUT("/update/:id", handler.UpdateUser)
	userRouter.DELETE("/delete/:id", handler.DeleteUser)

	// Create subRouter for Courses
	courseRouter := mainRouter.Group("/courses")

	courseRouter.POST("/create", handler.CreateCourse)
	courseRouter.GET("/getall", handler.GetCourses)
	courseRouter.GET("/:id", handler.GetCourseById)
	courseRouter.GET("/:id/lessons", handler.GetLessonsbyCourse)
	courseRouter.GET("/:id/enrollments", handler.GetEnrolledUsersByCourse)
	courseRouter.GET("/popular", handler.GetMostPopularCourses)
	courseRouter.PUT("/update/:id", handler.UpdateCourse)
	courseRouter.DELETE("/delete/:id", handler.DeleteCourse)

	// Create subRouter for Lessons
	lessonRouter := mainRouter.Group("/lessons")

	lessonRouter.POST("/create", handler.CreateLesson)
	lessonRouter.GET("/getall", handler.GetLessons)
	lessonRouter.GET("/:id", handler.GetLessonById)
	lessonRouter.PUT("/update/:id", handler.UpdateLesson)
	lessonRouter.DELETE("/delete/:id", handler.DeleteLesson)

	// Create subRouter for Enrollments
	enrollmentRouter := mainRouter.Group("/enrollments")

	enrollmentRouter.POST("/create", handler.CreateEnrollment)
	enrollmentRouter.GET("/getall", handler.GetEnrollments)
	enrollmentRouter.GET("/:id", handler.GetEnrollmentById)
	enrollmentRouter.PUT("/update/:id", handler.UpdateEnrollment)
	enrollmentRouter.DELETE("/delete/:id", handler.DeleteEnrollment)

	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
