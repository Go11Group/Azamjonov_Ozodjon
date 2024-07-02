package handler

import (
	"fmt"
	"language_learning_app/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create
func (h *Handler) CreateEnrollment(c *gin.Context) {
	newEnrollment := model.Enrollment{}

	// take json from request body to new struct
	err := c.BindJSON(&newEnrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting json from body: %s", err),
		})
		log.Printf("Error with getting json from URL body: %s\n", err)
		return
	}

	// use courses struct's method to create
	err = h.Enrollment.CreateEnrollment(&newEnrollment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with create new enrollment into database: %s", err),
		})
		log.Printf("Error with create new enrollment into database: %s\n", err)
		return
	}

	// send message to response 'there is no error'
	c.JSON(http.StatusOK, "Created new enrollment successfully")
}

// Read all
func (h *Handler) GetEnrollments(c *gin.Context) {
	filter := model.EnrollmentFilter{}

	// get filter parametres from URL query
	UserId, hasUserId := c.GetQuery("user_id")
	if hasUserId {
		filter.UserId = &UserId
	}
	CourseId, hasCourseId := c.GetQuery("course_id")
	if hasCourseId {
		filter.CourseId = &CourseId
	}
	EnrollmentDate, hasEnrollmentDate := c.GetQuery("enrollment_date")
	if hasEnrollmentDate {
		enrollmentdate, err := time.Parse("2006-01-02", EnrollmentDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "StatusInternalServerError",
				"message": fmt.Sprintf("Error with Parse string to time: %s", err),
			})
			log.Printf("Error with Parse string to time: %s\n", err)
			return
		}
		filter.EnrollmentDate = &enrollmentdate
	}
	limit, hasLimit := c.GetQuery("limit")
	if hasLimit {
		filter.Limit = &limit
	}
	offset, hasOffset := c.GetQuery("offset")
	if hasOffset {
		filter.Offset = &offset
	}

	// use enrollment struct's mothod to get data from database
	enrollments, err := h.Enrollment.GetEnrollments(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with getting enrollments from database: %s", err),
		})
		log.Printf("Error with getting enrollments from database: %s\n", err)
		return
	}

	// send data in response body
	c.JSON(http.StatusOK, enrollments)
}

// Read by id
func (h *Handler) GetEnrollmentById(c *gin.Context) {
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

	// use enrollments struct's method to get data by id
	enrollment, err := h.Enrollment.GetEnrollmentById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with getting enrollment by id from database: %s", err),
		})
		log.Printf("Error with getting enrollment by id from database: %s\n", err)
		return
	}

	// send data in response body
	c.JSON(http.StatusOK, enrollment)
}

// Update
func (h *Handler) UpdateEnrollment(c *gin.Context) {
	enrollment := model.Enrollment{}

	// get json from request body
	err := c.BindJSON(&enrollment)
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

	// put id to enrollments's CoursesId field
	enrollment.CourseId = id

	// use enrollments struct's method to update the data
	err = h.Enrollment.UpdateEnrollment(&enrollment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with updating enrollment in database: %s", err),
		})
		log.Printf("Error with updating enrollment in database: %s\n", err)
		return
	}

	// send message to response 'there is no error'
	c.JSON(http.StatusOK, "enrollment was updated successfully")
}

// Delete
func (h *Handler) DeleteEnrollment(c *gin.Context) {
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

	// use enrollment struct's method to delete the data
	err := h.Enrollment.DeleteEnrollment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with deleting enrollment in database: %s", err),
		})
		log.Printf("Error with deleting enrollment in database: %s\n", err)
		return
	}

	// send message to response 'there is no error'
	c.JSON(http.StatusOK, "enrollment was deleted successfully")
}
