package api

import (
	"auth_service/api/handler"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Router(handler *handler.Handler, db *sql.DB) *gin.Engine {
	router := gin.Default()

	//authMiddleware := middleware.JWTMiddleware()

	//r := router.Group("/api/v1").Use(authMiddleware)

	router.POST("/auth/register", handler.Register)
	router.POST("/auth/login", handler.Login)
	//router.POST("/auth/logout", handler.Logout)
	router.POST("/auth/refresh", handler.RefreshToken)
	//router.POST("/auth/reset-password", handler.ResetPassword)

	return router
}
