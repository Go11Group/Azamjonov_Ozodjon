package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imtihon/model"
	"net/http"
)

// Enrollment retrieves an enrollment by ID.
func (h *Handler) Enrollment(c *gin.Context) {
	id := c.Param("id")
	enrollment, err := h.Enrollments.EnrollmentGet(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrollment)
}

// EnrollmentCreate creates a new enrollment.
func (h *Handler) EnrollmentCreate(c *gin.Context) {
	var enrollment model.Enrollments
	if err := c.BindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Enrollments.EnrollmentCreate(enrollment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create enrollment"})
		return
	}

	c.JSON(http.StatusCreated, enrollment)
}

// EnrollmentUpdate updates an existing enrollment by ID.
func (h *Handler) EnrollmentUpdate(c *gin.Context) {
	id := c.Param("id")
	var enrollment model.Enrollments
	if err := c.BindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	enrollment.EnrollmentId = id
	if err := h.Enrollments.EnrollmentUpdate(enrollment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrollment)
}

// EnrollmentDelete deletes an enrollment by ID.
func (h *Handler) EnrollmentDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Enrollments.EnrollmentDelete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// EnrollmentGet retrieves enrollments based on query parameters.
func (h *Handler) EnrollmentGet(c *gin.Context) {
	params := make(map[string]interface{})
	var arr []interface{}
	var limit, offset string

	query := "SELECT * FROM enrollments WHERE deleted_at = 0"

	// Apply filters based on query parameters
	if enrollmentID := c.Query("enrollmentid"); enrollmentID != "" {
		params["enrollmentid"] = enrollmentID
		query += " AND enrollmentid = :enrollmentid"
	}
	if userID := c.Query("userid"); userID != "" {
		params["userid"] = userID
		query += " AND userid = :userid"
	}
	if courseID := c.Query("courseid"); courseID != "" {
		params["courseid"] = courseID
		query += " AND courseid = :courseid"
	}
	if enrollmentDate := c.Query("enrollmentdate"); enrollmentDate != "" {
		params["enrollmentdate"] = enrollmentDate
		query += " AND enrollmentdate = :enrollmentdate"
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

	// Execute query to get enrollments from database
	enrollments, err := h.Enrollments.Get(query, arr)
	if err != nil {
		fmt.Println(err) // Print error to console
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}

// EnrollmentsByCourseId retrieves enrollments by course ID.
func (h *Handler) EnrollmentsByCourseId(c *gin.Context) {
	courseId := c.Param("course_id")

	// Retrieve enrollments by course ID from repository
	enrollments, err := h.Enrollments.EnrolledByCourseId(courseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course_id": courseId, "enrollments": enrollments})
}
