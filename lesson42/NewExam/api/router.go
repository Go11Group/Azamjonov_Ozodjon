package api

import (
	"github.com/gin-gonic/gin"
	"my_project/api/handlers"
	"my_project/strorage/postgres"
)

func RooterApi(course *postgres.CourseRepository, enrollment *postgres.EnrollmentRepository, lesson *postgres.LessonRepository, user *postgres.UserRepository) *gin.Engine {
	gn := gin.Default()
	handler := handlers.NewHandler(user, course, lesson, enrollment)
	routeUser := gn.Group("/api/user")
	routeUser.POST("/create", handler.CreateUser)
	routeUser.GET("/get", handler.GetUser)
	routeUser.GET("/courses/:id", handler.GetCourseByUserId)
	routeUser.GET("/id/:id", handler.GetCourseById)
	routeUser.GET("/popular/", handler.ShowPopularCourse)
	routeUser.GET("/get/:name/:email", handler.GetUserByEmailOrName)
	routeUser.PUT("/update/:id", handler.UpdateUser)
	routeUser.DELETE("/delete/:id", handler.DeleteUser)

	// Course router
	routeCourse := gn.Group("/api/course")

	routeCourse.POST("/create", handler.CreateCourse)
	routeCourse.GET("/get/", handler.GetCourse)
	routeCourse.GET("/id/:id", handler.GetCourseById)
	routeCourse.GET("/lessons/:id", handler.GetLessonsByCourseId)
	routeCourse.GET("/enrollments/:id", handler.GetUserWithEnrollmentByCourseId)
	routeCourse.PUT("/update/:id", handler.UpdateCourse)
	routeCourse.DELETE("/delete/:id", handler.DeleteCourse)

	// Lessons router
	routeLesson := gn.Group("/api/lesson")

	routeLesson.POST("/create", handler.CreateLesson)
	routeLesson.GET("/get", handler.GetLesson)
	routeLesson.GET("/id/:id", handler.GetLessonById)
	routeLesson.PUT("/update/:id", handler.UpdateLesson)
	routeLesson.DELETE("/delete/:id", handler.DeleteLesson)
	//

	// enrollment router
	routeEnrollment := gn.Group("/api/enrollment")
	routeEnrollment.POST("/create", handler.CreateEnrollment)
	routeEnrollment.GET("/get", handler.GetEnrollment)
	routeEnrollment.GET("/id/:id", handler.GetEnrollmentById)
	routeEnrollment.PUT("/update/:id", handler.UpdateEnrollment)
	routeEnrollment.DELETE("/delete/:id", handler.DeleteEnrollment)
	return gn
}
