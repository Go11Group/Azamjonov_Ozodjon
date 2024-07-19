package handler

import (
	"context"
	"net/http"
	"strings"

	pb "auth_service/generated/genproto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(ctx *gin.Context) {
	var req pb.RegisterUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error binding JSON:", err)
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	resp, err := h.User.RegisterUser(&req)
	if err != nil {
		h.Logger.Error("Error registering user:", err)
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(200, resp)
}

func (h *Handler) Login(ctx *gin.Context) {
	var req pb.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.Logger.Error("Error binding JSON:", err)
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := h.User.LoginUser(&req)
	if err != nil {
		h.Logger.Error("Error logging in user:", err)
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(200, resp)
}

func (h *Handler) GetProfile(ctx *gin.Context) {
	// Assuming you get user ID from middleware
	userID := ctx.GetString("userID")
	req := pb.ById{Id: userID}

	res, err := h.User.GetUser(&req)
	if err != nil {
		h.Logger.Error("Failed to get user by id", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) UpdateProfile(ctx *gin.Context) {
	var req pb.UpdateUserReq
	if err := ctx.BindJSON(&req); err != nil {
		h.Logger.Error("Failed to bind request", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.UpdateUser(&req)
	if err != nil {
		h.Logger.Error("Failed to update user", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) GetUsers(ctx *gin.Context) {
	var req pb.PageLimit
	if err := ctx.BindQuery(&req); err != nil {
		h.Logger.Error("Failed to bind request", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.GetAllUsers(&req)
	if err != nil {
		h.Logger.Error("Failed to get all users", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("user_id")
	req := pb.ById{Id: id}

	_, err := h.User.DeleteUser(&req)
	if err != nil {
		h.Logger.Error("Failed to delete user", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
}

func (h *Handler) ResetPassword(ctx *gin.Context) {
	var req pb.ByEmail
	if err := ctx.BindJSON(&req); err != nil {
		h.Logger.Error("Failed to bind request", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.User.ResetPassword(&req)
	if err != nil {
		h.Logger.Error("Failed to reset password", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password reset instructions sent to your email"})
}

func (h *Handler) RefreshToken(ctx *gin.Context) {
	var req pb.RefreshTokenReq
	if err := ctx.BindJSON(&req); err != nil {
		h.Logger.Error("Failed to bind request", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.RefreshToken(&req)
	if err != nil {
		h.Logger.Error("Failed to refresh token", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	req := &pb.LogoutReq{AccessToken: token}

	res, err := h.User.Logout(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetEcoPoints(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	req := pb.ById{Id: userID}

	res, err := h.User.GetEcoPoints(&req)
	if err != nil {
		h.Logger.Error("Failed to get eco-points", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) AddEcoPoints(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	var req pb.AddEcoPointsReq
	if err := ctx.BindJSON(&req); err != nil {
		h.Logger.Error("Failed to bind request", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserId = userID

	res, err := h.User.AddEcoPoints(&req)
	if err != nil {
		h.Logger.Error("Failed to add eco-points", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) GetEcoPointsHistory(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	var req pb.PageLimit1
	if err := ctx.BindQuery(&req); err != nil {
		h.Logger.Error("Failed to bind request", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Userid = userID

	res, err := h.User.GetEcoPointsHistory(&req)
	if err != nil {
		h.Logger.Error("Failed to get eco-points history", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
