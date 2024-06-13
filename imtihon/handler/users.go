package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/imtihon/model"
	"net/http"
)

func (h *Handler) User(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Users.User_Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) UserCreate(c *gin.Context) {
	var user model.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You can set default values or perform validation here if needed

	if err := h.Users.User_Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *Handler) UserUpdate(c *gin.Context) {
	id := c.Param("id")
	var user model.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = id
	if err := h.Users.User_Update(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (h *Handler) UserDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.Users.User_Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
