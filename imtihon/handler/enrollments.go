package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/imtihon/model"
	"net/http"
)

func (h *Handler) Enrollment(c *gin.Context) {
	id := c.Param("id")
	enrollment, err := h.Enrollments.Enrollment_Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrollment)
}

func (h *Handler) EnrollmentCreate(c *gin.Context) {
	var enrollment model.Enrollments
	if err := c.BindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Enrollments.Enrollment_Create(enrollment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, enrollment)
}

func (h *Handler) EnrollmentUpdate(c *gin.Context) {
	id := c.Param("id")
	var enrollment model.Enrollments
	if err := c.BindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	enrollment.EnrollmentId = id
	if err := h.Enrollments.Enrollment_Update(enrollment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrollment)
}

func (h *Handler) EnrollmentDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Enrollments.Enrollment_Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
