package handler

import (
	"atto/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateCard(ctx *gin.Context) {
	var card models.CreateCard
	if err := ctx.ShouldBindJSON(&card); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Card.Create(&card); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (h *Handler) GetCard(ctx *gin.Context) {
	id := ctx.Param("id")
	card, err := h.Card.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "card not found"})
		return
	}

	ctx.JSON(http.StatusOK, card)
}

func (h *Handler) UpdateCard(ctx *gin.Context) {
	id := ctx.Param("id")
	var card models.Card
	if err := ctx.ShouldBindJSON(&card); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	card.Id = id

	if err := h.Card.Update(&card); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteCard(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.Card.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h *Handler) GetAllCards(ctx *gin.Context) {
	cards, err := h.Card.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cards)
}
