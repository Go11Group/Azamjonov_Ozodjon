package handler

import (
	"encoding/json"
	"fmt"
	"lesson62/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (h *Handler) CreateRestaurant(ctx *gin.Context) {
	newres := models.RestaurantCreate{}
	// "a1ebb162-92cb-4f72-b9b2-141f26031575"
	err := json.NewDecoder(ctx.Request.Body).Decode(&newres)
	if err != nil {
		h.Logger.Error("Error with decoding url body", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with decoding url body: %s", err))
		return
	}

	reswithFullInfo := models.RestaurantGet{
		Id:          uuid.NewString(),
		Title:       newres.Title,
		Description: newres.Description,
		Address:     newres.Address,
		ProneNumber: newres.ProneNumber,
		Time: models.Time{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	req, err := json.Marshal(reswithFullInfo)
	if err != nil {
		h.Logger.Error("Error with marshel restaulrant info", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with marshel restaulrant info: %s", err))
		return
	}

	err = h.Restaurant.Create(ctx, reswithFullInfo.Id, string(req))
	if err != nil {
		h.Logger.Error("Error with creating restaurant", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with creating restaurant: %s", err))
		return
	}

	ctx.JSON(http.StatusOK, reswithFullInfo)
}

func (h *Handler) GetAllRestaurants(ctx *gin.Context) {
	resp, err := h.Restaurant.GetAll(ctx)
	if err != nil {
		h.Logger.Error("Error with Getting all restaurant", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with Getting all restaurant: %s", err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.Logger.Error("Error with Getting id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			"Error with Getting id")
		return
	}

	resp, err := h.Restaurant.GetById(ctx, id)
	if err != nil {
		h.Logger.Error("Error with Getting data from database", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with Getting data from database: %s", err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateRestaurant(ctx *gin.Context) {
	restaurantToUpdate := models.RestaurantUpdate{}
	// "a1ebb162-92cb-4f72-b9b2-141f26031575"
	err := json.NewDecoder(ctx.Request.Body).Decode(&restaurantToUpdate)
	if err != nil {
		h.Logger.Error("Error with decoding url body", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with decoding url body: %s", err))
		return
	}

	err = h.Restaurant.Update(ctx, &restaurantToUpdate)
	if err != nil {
		h.Logger.Error("Error with updating restaurant", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with updating restaurant: %s", err))
		return
	}
	ctx.JSON(http.StatusOK, "restaurant was updated successfully")
}

func (h *Handler) DeleteRestaurant(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.Logger.Error("Error with Getting id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			"Error with Getting id")
		return
	}

	err := h.Restaurant.Delete(ctx, id)
	if err != nil {
		h.Logger.Error("Error with Getting data from database", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with Getting data from database: %s", err))
		return
	}
	ctx.JSON(http.StatusOK, "restaurant was deleted successfully")
}
