package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imtihon/model"
	"net/http"
)

// Lesson retrieves a lesson by ID.
func (h *Handler) Lesson(c *gin.Context) {
	id := c.Param("id")
	lesson, err := h.Lessons.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lesson)
}

// LessonCreate creates a new lesson.
func (h *Handler) LessonCreate(c *gin.Context) {
	var lesson model.Lessons
	if err := c.BindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Lessons.Create(lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, lesson)
}

// LessonUpdate updates an existing lesson by ID.
func (h *Handler) LessonUpdate(c *gin.Context) {
	id := c.Param("id")
	var lesson model.Lessons
	err := c.ShouldBindJSON(&lesson)
	if err != nil {
		fmt.Println("hello")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
	}
	if err := h.Lessons.Update(id, lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lesson)
}

// LessonDelete deletes a lesson by ID.
func (h *Handler) LessonDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Lessons.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// LessonGet retrieves lessons based on query parameters.
func (h *Handler) LessonGet(c *gin.Context) {
	params := make(map[string]interface{})
	var arr []interface{}
	var limit, offset string

	query := "SELECT * FROM lessons WHERE deleted_at = 0"

	// Apply filters based on query parameters
	if lessonID := c.Query("lessonid"); lessonID != "" {
		params["lessonid"] = lessonID
		query += " AND lessonid = :lessonid"
	}
	if courseID := c.Query("courseid"); courseID != "" {
		params["courseid"] = courseID
		query += " AND courseid = :courseid"
	}
	if title := c.Query("title"); title != "" {
		params["title"] = title
		query += " AND title = :title"
	}
	if content := c.Query("content"); content != "" {
		params["content"] = content
		query += " AND content = :content"
	}

	// Apply limit and offset if specified
	if lim := c.Query("limit"); lim != "" {
		params["limit"] = lim
		limit = " LIMIT :limit"
	}
	if off := c.Query("offset"); off != "" {
		params["offset"] = off
		offset = " OFFSET :offset"
	}

	// Construct final query with optional limit and offset
	query = query + limit + offset
	query, arr = ReplaceQueryParamsU(query, params)

	// Execute query to get lessons from database
	lessons, err := h.Lessons.Get(query, arr)
	if err != nil {
		fmt.Println(err) // Print error to console
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lessons)
}

// LessonsByCourseId retrieves lessons by course ID.
func (h *Handler) LessonsByCourseId(c *gin.Context) {
	courseId := c.Param("course_id")

	// Retrieve lessons by course ID from repository
	lessons, err := h.Lessons.LessonsByCourseId(courseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course_id": courseId, "lessons": lessons})
}
