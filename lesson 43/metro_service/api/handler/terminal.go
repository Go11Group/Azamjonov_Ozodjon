package handler

import (
	"atto/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateTerminal(ctx *gin.Context) {
	var terminal models.CreateTerminal
	if err := ctx.ShouldBindJSON(&terminal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Terminal.Create(&terminal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (h *Handler) GetTerminal(ctx *gin.Context) {
	id := ctx.Param("id")
	terminal, err := h.Terminal.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "terminal not found"})
		return
	}

	ctx.JSON(http.StatusOK, terminal)
}

func (h *Handler) UpdateTerminal(ctx *gin.Context) {
	id := ctx.Param("id")
	var terminal models.Terminal
	if err := ctx.ShouldBindJSON(&terminal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	terminal.Id = id

	if err := h.Terminal.Update(&terminal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteTerminal(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.Terminal.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h *Handler) GetAllTerminals(ctx *gin.Context) {
	terminals, err := h.Terminal.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, terminals)
}
