package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/imtihon/model"
	"net/http"
)

func (h *Handler) Course(c *gin.Context) {
	id := c.Param("id")
	course, err := h.Courses.Course_Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *Handler) CourseCreate(c *gin.Context) {
	var course model.Courses
	if err := c.BindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Courses.Course_Create(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, course)
}

func (h *Handler) CourseUpdate(c *gin.Context) {
	id := c.Param("id")
	var course model.Courses
	if err := c.BindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course.Id = id
	if err := h.Courses.Course_Update(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *Handler) CourseDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Courses.Course_Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
