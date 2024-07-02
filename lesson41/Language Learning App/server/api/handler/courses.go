package handler

import (
	"fmt"
	"language_learning_app/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create
func (h *Handler) CreateCourse(c *gin.Context) {
	newCourse := model.Course{}

	// take json from request body to new struct
	err := c.BindJSON(&newCourse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting json from body: %s", err),
		})
		log.Printf("Error with getting json from courses URL body: %s\n", err)
		return
	}

	// use courses struct's method to create
	err = h.Courses.CreateCourse(&newCourse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with create new course into database: %s", err),
		})
		log.Printf("Error with create course into databasee: %s\n", err)
		return
	}

	// send message to response 'there is no error'
	c.JSON(http.StatusOK, "Created new Course successfully")
}

// Read all
func (h *Handler) GetCourses(c *gin.Context) {
	filter := model.CourseFilter{}

	// get filter parametres from URL query
	title, hasTitle := c.GetQuery("title")
	if hasTitle {
		filter.Title = &title
	}
	description, hasDescription := c.GetQuery("description")
	if hasDescription {
		filter.Description = &description
	}
	limit, hasLimit := c.GetQuery("limit")
	if hasLimit {
		filter.Limit = &limit
	}
	offset, hasOffset := c.GetQuery("offset")
	if hasOffset {
		filter.Offset = &offset
	}

	// use courses struct's mothod to get data from database
	courses, err := h.Courses.GetCourses(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with getting courses from database%s\n", err),
		})
		log.Printf("Error with getting courses from database: %s\n", err)
		return
	}

	// send data in response body
	c.JSON(http.StatusOK, courses)
}

// Read by id
func (h *Handler) GetCourseById(c *gin.Context) {
	// get id from URL params
	id, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: id not found in URL params",
		})
		log.Printf("Error: id not found in URL params")
		return
	}

	// use courses struct's method to get data by id
	course, err := h.Courses.GetCourseById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with getting course by id from database: %s", err),
		})
		log.Printf("Error with getting course by id from database: %s\n", err)
		return
	}

	// send data in response body
	c.JSON(http.StatusOK, course)
}

func (h *Handler) GetLessonsbyCourse(c *gin.Context) {
	courseId, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: id not found in URL params",
		})
		log.Printf("Error: id not found in URL params")
		return
	}

	lessons, err := h.Courses.GetLessonsbyCourse(courseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with getting course's lessons from database: %s", err),
		})
		log.Printf("Error with getting course's lessons from database: %s\n", err)
		return
	}
	c.JSON(http.StatusOK, lessons)
}

func (h *Handler) GetEnrolledUsersByCourse(c *gin.Context) {
	courseId, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: id not found in URL params",
		})
		log.Printf("Error: id not found in URL params")
		return
	}

	users, err := h.Courses.GetEnrolledUsersByCourse(courseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with getting course's users from database: %s", err),
		})
		log.Printf("Error with getting course's users from database: %s\n", err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetMostPopularCourses(c *gin.Context) {
	startDate, hasit := c.GetQuery("start_date")
	if !hasit {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: start_date not found in URL params",
		})
		log.Printf("Error: start_date not found in URL params")
		return
	}
	endDate, hasit := c.GetQuery("end_date")
	if !hasit {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: end_date not found in URL params",
		})
		log.Printf("Error: end_date not found in URL params")
		return
	}
	limit, hasit := c.GetQuery("limit")
	if !hasit {
		limit = "5"
	}
	courses, err := h.Courses.GetMostPopularCourses(startDate, endDate, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with getting popular courses from database: %s", err),
		})
		log.Printf("Error with getting popular courses from database: %s\n", err)
		return
	}
	c.JSON(http.StatusOK, courses)
}

// Update
func (h *Handler) UpdateCourse(c *gin.Context) {
	course := model.Course{}

	// get json from request body
	err := c.BindJSON(&course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting json from Request body: %s", err),
		})
		log.Printf("Error with getting json from Request body: %s\n", err)
		return
	}

	// get id from URL params
	id, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: id not found in URL params",
		})
		log.Printf("Error: id not found in URL params")
		return
	}

	// put id to courses's CoursesId field
	course.CourseId = id

	// use courses struct's method to update the data
	err = h.Courses.UpdateCourse(&course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with updating course in database: %s", err),
		})
		log.Printf("Error with updating course in database: %s\n", err)
		return
	}

	// send message to response 'there is no error'
	c.JSON(http.StatusOK, "Course was updated successfully")
}

// Delete
func (h *Handler) DeleteCourse(c *gin.Context) {
	// get id from URL params
	id, hasId := c.Params.Get("id")
	if !hasId {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": "Error: id not found in URL params",
		})
		log.Printf("Error: id not found in URL params")
		return
	}

	// use courses struct's method to delete the data
	err := h.Courses.DeleteCourse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with deleting course in database: %s", err),
		})
		log.Printf("Error with deleting course in database: %s\n", err)
		return
	}

	// send message to response 'there is no error'
	c.JSON(http.StatusOK, "Course was deleted successfully")
}
