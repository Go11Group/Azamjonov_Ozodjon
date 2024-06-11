package handler

import (
	"net/http"
	"strconv"

	"github.com/Azamjonov_Ozodjon/lesson36/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Userr(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.User.User_Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UserCreate(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	if err := h.User.User_Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *Handler) UserUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updatedUser model.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	if err := h.User.User_Update(id, updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h *Handler) UserDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.User.User_Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
