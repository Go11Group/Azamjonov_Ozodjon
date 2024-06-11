package handler

import (
	"github.com/Azamjonov_Ozodjon/lesson36/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) Problem1(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid problem ID"})
		return
	}

	problem, err := h.Problem.Problem_Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting problem"})
		return
	}

	c.JSON(http.StatusOK, problem)
}

func (h *Handler) ProblemCreate(c *gin.Context) {
	var problem model.Problem
	if err := c.ShouldBindJSON(&problem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	if err := h.Problem.Problem_Create(problem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating problem"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Problem created successfully"})
}

func (h *Handler) ProblemUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid problem ID"})
		return
	}

	var updatedProblem model.Problem
	if err := c.ShouldBindJSON(&updatedProblem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	if err := h.Problem.Problem_Update(id, updatedProblem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating problem"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Problem updated successfully"})
}

func (h *Handler) ProblemDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid problem ID"})
		return
	}

	if err := h.Problem.Problem_Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting problem"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Problem deleted successfully"})
}
