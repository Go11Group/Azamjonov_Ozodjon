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

func (h *Handler) CreateUser(ctx *gin.Context) {
	newres := models.UserCreate{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&newres)
	if err != nil {
		h.Logger.Error("Error with decoding url body", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with decoding url body: %s", err))
		return
	}

	reswithFullInfo := models.UserGet{
		Id:        uuid.NewString(),
		Firstname: newres.Firstname,
		Lastname:  newres.Lastname,
		Age:       newres.Age,
		Role:      newres.Role,
		Time: models.Time{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	req, err := json.Marshal(reswithFullInfo)
	if err != nil {
		h.Logger.Error("Error with marshel user info", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with marshel user info: %s", err))
		return
	}

	err = h.Users.Create(ctx, reswithFullInfo.Id, string(req))
	if err != nil {
		h.Logger.Error("Error with creating user", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with creating user: %s", err))
		return
	}

	ctx.JSON(http.StatusOK, reswithFullInfo)
}

func (h *Handler) GetAllUsers(ctx *gin.Context) {
	resp, err := h.Restaurant.GetAll(ctx)
	if err != nil {
		h.Logger.Error("Error with Getting all users", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with Getting all users: %s", err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		h.Logger.Error("Error with Getting id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			"Error with Getting id")
		return
	}

	resp, err := h.Users.GetById(ctx, id)
	if err != nil {
		h.Logger.Error("Error with Getting data from database", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with Getting data from database: %s", err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	userToUpdate := models.UserUpdate{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&userToUpdate)
	if err != nil {
		h.Logger.Error("Error with decoding url body", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with decoding url body: %s", err))
		return
	}

	err = h.Users.Update(ctx, &userToUpdate)
	if err != nil {
		h.Logger.Error("Error with updating user", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			fmt.Sprintf("Error with updating user: %s", err))
		return
	}
	ctx.JSON(http.StatusOK, "user was updated successfully")
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, "user was deleted successfully")
}
