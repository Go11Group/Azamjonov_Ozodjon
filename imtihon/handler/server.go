package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handlerr(h *Handler) *http.Server {
	r := gin.Default()

	//user

	//http://localhost:8080/user
	r.GET("/user/:id", h.UserGetById)
	//http://localhost:8080/user/get
	r.GET("/user/get", h.UserGet)
	//http://localhost:8080/user
	//{
	//	"name": "John Doe",
	//	"email": "john.doe@example.com",
	//	"birthday": "1990-01-01",
	//	"password": "johnspassword"
	//}

	r.POST("/user", h.UserCreate)
	//http://localhost:8080/user/39630973-99c1-43bf-9ae6-66f2fc6ffd66

	r.PUT("/user/:id", h.UserUpdate)
	//http://localhost:8080/user/39630973-99c1-43bf-9ae6-66f2fc6ffd66
	r.DELETE("/user/:id", h.UserDelete)

	//course
	//http://localhost:8080/course/350d70c5-3481-4811-9c81-baf1cf63025e
	r.GET("/course/:id", h.CourseGetById)
	//http://localhost:8080/course/get
	r.GET("/course/get", h.CourseGet)
	//http://localhost:8080/course
	r.POST("/course", h.CourseCreate)
	//http://localhost:8080/course/f9649b79-dfc2-4900-bf17-f65a75985758
	r.PUT("/course/:id", h.CourseUpdate)
	//http://localhost:8080/course/c93e0913-56e8-40d8-bb8b-92aacef1a072
	r.DELETE("/course/:id", h.CourseDelete)

	//lesson
	//http://localhost:8080/lesson/11d94428-7751-4cce-a7cc-ad4f70f8309c
	r.GET("/lesson/:id", h.Lesson)
	//http://localhost:8080/lesson/get?limit=2
	r.GET("/lesson/get", h.LessonGet)
	//http://localhost:8080/lesson
	//{
	//    "course_id": "1bbfc55a-8897-434f-906c-66224718c047",
	//    "title": "Lesson 1: Basics of Go",
	//    "content": "Introduction to variables, loops, and functions in Go."
	//}
	r.POST("/lesson", h.LessonCreate)
	//http://localhost:8080/lesson/1bbfc55a-8897-434f-906c-66224718c047
	//{
	//    "title": "Lesson 1: Basics of Go Updated",
	//    "content": "Updated content for introduction to Go.",
	//    "updated_at": "2024-06-14T00:00:00Z",
	//    "deleted_at": 0
	//}
	r.PUT("/lesson/:id", h.LessonUpdate)
	//http://localhost:8080/lesson/00000000-0000-0000-0000-000000000000
	//{
	//    "title": "Lesson 1: Basics of Go Updated",
	//    "content": "Updated content for introduction to Go.",
	//    "updated_at": "2024-06-14T00:00:00Z",
	//    "deleted_at": 0
	//}
	r.DELETE("/lesson/:id", h.LessonDelete)

	//enrollment
	//localhost:8080/enrollment/1a9f921a-505a-44d7-b78b-44f13841f82a
	r.GET("/enrollment/:id", h.Enrollment)
	//localhost:8080/enrollment/get
	r.GET("/enrollment/get", h.EnrollmentGet)
	//localhost:8080/enrollment
	//{
	//"user_id": "958b2e24-ef28-4f36-bc42-4d9b0c716024",
	//"course_id": "350d70c5-3481-4811-9c81-baf1cf63025e",
	//"enrollment_date": "2024-06-13T17:55:39"
	//}
	r.POST("/enrollment", h.EnrollmentCreate)
	//localhost:8080/enrollment/1a9f921a-505a-44d7-b78b-44f13841f82a
	r.PUT("/enrollment/:id", h.EnrollmentUpdate)
	//localhost:8080/enrollment/1a9f921a-505a-44d7-b78b-44f13841f82a
	r.DELETE("/enrollment/:id", h.EnrollmentDelete)

	//http://localhost:8080/users/958b2e24-ef28-4f36-bc42-4d9b0c716024/courses
	r.GET("/users/:user_id/courses", h.GetCoursesByUserID)
	//http://localhost:8080/courses/350d70c5-3481-4811-9c81-baf1cf63025e/lessons
	r.GET("/courses/:course_id/lessons", h.LessonsByCourseId)
	//http://localhost:8080/courses/350d70c5-3481-4811-9c81-baf1cf63025e/enrollments
	r.GET("/courses/:course_id/enrollments", h.EnrollmentsByCourseId)
	//http://localhost:8080/users/search?name=Jane
	r.GET("/users/search", h.UserSearch)
	//http://localhost:8080/courses/popular?start_date=2000-01-01&end_date=2026-12-31
	r.GET("/courses/popular", h.GetPopularCourses)
	return &http.Server{Addr: ":8080", Handler: r}
}
