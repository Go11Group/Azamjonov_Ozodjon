// handler/lessons.go

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/imtihon/model"
	"net/http"
)

func (h *Handler) Lesson(c *gin.Context) {
	id := c.Param("id")
	lesson, err := h.Lessons.Lesson_Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lesson)
}

func (h *Handler) LessonCreate(c *gin.Context) {
	var lesson model.Lessons
	if err := c.BindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Lessons.Lesson_Create(lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, lesson)
}

func (h *Handler) LessonUpdate(c *gin.Context) {
	id := c.Param("id")
	var lesson model.Lessons
	if err := c.BindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lesson.LessonId = id
	if err := h.Lessons.Lesson_Update(lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lesson)
}

func (h *Handler) LessonDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Lessons.Lesson_Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
