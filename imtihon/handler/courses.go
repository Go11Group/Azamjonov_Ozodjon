package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imtihon/model"
	"net/http"
	"time"
)

// CourseGet retrieves a course by its ID.
func (h *Handler) CourseGetById(c *gin.Context) {
	id := c.Param("id")
	course, err := h.Courses.CourseGet(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}

// CourseCreate creates a new course.
func (h *Handler) CourseCreate(c *gin.Context) {
	var course model.Courses
	if err := c.BindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()
	if err := h.Courses.CourseCreate(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, course)
}

// CourseUpdate updates an existing course by its ID.
func (h *Handler) CourseUpdate(c *gin.Context) {
	id := c.Param("id")
	var course model.Courses
	if err := c.BindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course.Id = id
	if err := h.Courses.CourseUpdate(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}

// CourseDelete deletes a course by its ID.
func (h *Handler) CourseDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Courses.CourseDelete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// GetCoursesByUserID retrieves courses associated with a specific user ID.
func (h *Handler) GetCoursesByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	courses, err := h.Courses.GetCoursesByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_id": userID, "courses": courses})
}

// GetPopularCourses retrieves courses that are most popular within a specified time period.
func (h *Handler) GetPopularCourses(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	// Check if start_date and end_date are provided
	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	// Retrieve popular courses within the specified time period
	courses, err := h.Courses.GetPopularCourses(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get popular courses"})
		return
	}

	// Construct response with time_period and popular_courses
	response := gin.H{
		"time_period": gin.H{
			"start_date": startDate,
			"end_date":   endDate,
		},
		"popular_courses": courses,
	}
	c.JSON(http.StatusOK, response)
}

// CourseGet retrieves courses based on query parameters.
func (h *Handler) CourseGet(c *gin.Context) {
	params := make(map[string]interface{})
	var arr []interface{}
	var limit, offset string

	query := `select * from courses where deleted_at=0`
	// Apply filters based on query parameters
	if title := c.Query("title"); title != "" {
		params["title"] = title
		query += " and title = :title"
	}
	if description := c.Query("description"); description != "" {
		params["description"] = description
		query += " and description = :description"
	}
	if lim := c.Query("limit"); lim != "" {
		params["limit"] = lim
		limit = ` LIMIT :limit`
	}
	if off := c.Query("offset"); off != "" {
		params["offset"] = off
		offset = ` OFFSET :offset`
	}

	// Concatenate limit and offset to the query
	query = query + limit + offset

	// Replace named parameters with positional parameters
	query, arr = ReplaceQueryParamsU(query, params)
	fmt.Println(query, arr)

	// Retrieve courses from the database with applied filters
	courses, err := h.Courses.Get(query, arr)
	if err != nil {
		fmt.Println(err) // Print error to console
		c.JSON(200, gin.H{"ERROR IN GET": err})
		return
	}

	c.JSON(200, courses)
}
