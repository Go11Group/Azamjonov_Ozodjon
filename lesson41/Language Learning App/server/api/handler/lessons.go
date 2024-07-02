package handler

import (
	"fmt"
	"language_learning_app/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create
func (h *Handler) CreateLesson(c *gin.Context) {
	newLesson := model.Lesson{}

	// take json from request body to new struct
	err := c.BindJSON(&newLesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with getting json from body: %s", err),
		})
		log.Printf("Error with getting json from URL body: %s\n", err)
		return
	}

	// use Lessons struct's method to create
	err = h.Lessons.CreateLesson(&newLesson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with create new lesson into database: %s", err),
		})
		log.Printf("Error with create new lesson into database: %s\n", err)
		return
	}

	// send message to response 'there is no error'
	c.JSON(http.StatusOK, "Created new Lesson successfully")
}

// Read all
func (h *Handler) GetLessons(c *gin.Context) {
	filter := model.LessonFilter{}

	// get filter parametres from URL query
	CourseId, hasCourseId := c.GetQuery("course_id")
	if hasCourseId {
		filter.CourseId = &CourseId
	}
	title, hasTitle := c.GetQuery("title")
	if hasTitle {
		filter.CourseId = &title
	}
	context, hasContent := c.GetQuery("content")
	if hasContent {
		filter.Content = &context
	}
	limit, hasLimit := c.GetQuery("limit")
	if hasLimit {
		filter.Limit = &limit
	}
	offset, hasOffset := c.GetQuery("offset")
	if hasOffset {
		filter.Offset = &offset
	}

	// use lessons struct's mothod to get data from database
	lessons, err := h.Lessons.GetLessons(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with getting lessons from database: %s", err),
		})
		log.Printf("Error with getting lessons from database: %s\n", err)
		return
	}

	// send data in response body
	c.JSON(http.StatusOK, lessons)
}

// Read by id
func (h *Handler) GetLessonById(c *gin.Context) {
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

	// use lessons struct's method to get data by id
	lesson, err := h.Lessons.GetLessonById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with getting lesson by id from database: %s", err),
		})
		log.Printf("Error with getting lesson by id from database: %s\n", err)
		return
	}

	// send data in response body
	c.JSON(http.StatusOK, lesson)
}

// Update
func (h *Handler) UpdateLesson(c *gin.Context) {
	lesson := model.Lesson{}

	// get json from request body
	err := c.BindJSON(&lesson)
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

	// put id to lessons's LessonId field
	lesson.LessonId = id

	// use lessons struct's method to update the data
	err = h.Lessons.UpdateLesson(&lesson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with updating lesson in database: %s", err),
		})
		log.Printf("Error with updating lesson in database: %s\n", err)
		return
	}

	// send message to response 'there is no error'
	c.JSON(http.StatusOK, "lesson was updated successfully")
}

// Delete
func (h *Handler) DeleteLesson(c *gin.Context) {
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

	// use lessons struct's method to delete the data
	err := h.Lessons.DeleteLesson(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "StatusInternalServerError",
			"message": fmt.Sprintf("Error with deleting lesson in database: %s", err),
		})
		log.Printf("Error with deleting lesson in database: %s\n", err)
		return
	}

	// send message to response 'there is no error'
	c.JSON(http.StatusOK, "lesson was deleted successfully")
}
