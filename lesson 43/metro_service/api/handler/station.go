package handler

import (
	"atto/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateStation(ctx *gin.Context) {
	var station models.CreateStation
	if err := ctx.ShouldBindJSON(&station); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Station.Create(&station); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (h *Handler) GetStation(ctx *gin.Context) {
	id := ctx.Param("id")
	station, err := h.Station.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "station not found"})
		return
	}

	ctx.JSON(http.StatusOK, station)
}

func (h *Handler) UpdateStation(ctx *gin.Context) {
	id := ctx.Param("id")
	var station models.Station
	if err := ctx.ShouldBindJSON(&station); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	station.Id = id

	if err := h.Station.Update(&station); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteStation(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.Station.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h *Handler) GetAllStations(ctx *gin.Context) {
	stations, err := h.Station.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, stations)
}
