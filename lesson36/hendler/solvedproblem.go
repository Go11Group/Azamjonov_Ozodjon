package handler

import (
	"net/http"
	"strconv"

	"github.com/Azamjonov_Ozodjon/lesson36/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SolvedProblemGet(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid solved problem ID"})
		return
	}

	solvedProblem, err := h.SolvedProblem.SolvedProblem_Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting solved problem"})
		return
	}

	c.JSON(http.StatusOK, solvedProblem)
}

func (h *Handler) SolvedProblemCreate(c *gin.Context) {
	var solvedProblem model.SolvedProblem
	if err := c.ShouldBindJSON(&solvedProblem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	if err := h.SolvedProblem.SolvedProblem_Create(solvedProblem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating solved problem"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Solved problem created successfully"})
}

func (h *Handler) SolvedProblemUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid solved problem ID"})
		return
	}

	var updatedSolvedProblem model.SolvedProblem
	if err := c.ShouldBindJSON(&updatedSolvedProblem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	if err := h.SolvedProblem.SolvedProblem_Update(id, updatedSolvedProblem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating solved problem"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Solved problem updated successfully"})
}

func (h *Handler) SolvedProblemDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid solved problem ID"})
		return
	}

	if err := h.SolvedProblem.SolvedProblem_Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting solved problem"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Solved problem deleted successfully"})
}
