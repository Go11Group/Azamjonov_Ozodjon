package handler

import (
	"atto/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateTransaction(ctx *gin.Context) {
	var transaction models.CreateTransaction
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Transaction.Create(&transaction); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (h *Handler) GetTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	transaction, err := h.Transaction.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (h *Handler) UpdateTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	var transaction models.Transaction
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction.Id = id

	if err := h.Transaction.Update(&transaction); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteTransaction(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.Transaction.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h *Handler) GetAllTransactions(ctx *gin.Context) {
	transactions, err := h.Transaction.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}
