package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/imtihon/3/Api_gateway/api/docs"
	hendler "github.com/imtihon/3/Api_gateway/api/handler"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter @title Eco swap
// @version 1.0
// @description API service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(h hendler.Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	// Item routes
	r.POST("/api/v1/items", h.CreateItem)
	r.PUT("/api/v1/items/:item_id", h.UpdateItem)
	r.DELETE("/api/v1/items/:item_id", h.DeleteItem)
	r.GET("/api/v1/itemes", h.GetAllItems)
	r.GET("/api/v1/items/:item_id", h.GetByIdItem)
	r.GET("/api/v1/items/search", h.SearchItemsAndFilter)

	// Swap routes
	r.POST("/api/v1/swaps", h.CreateChangeSwaps)
	r.PUT("/api/v1/swaps/:swap_id/accept", h.UpdateAcceptSwap)
	r.PUT("/api/v1/swaps/:swap_id/reject", h.UpdateRejectSwap)
	r.GET("/api/v1/swaps", h.GetChangedSwap)

	// Recycling center routes
	r.POST("/api/v1/recycling-centers", h.CreateAddRecyclingCenter)
	r.GET("/api/v1/recycling-centers", h.SearchRecyclingCenter)
	r.POST("/api/v1/recycling-submissions", h.CreateRecyclingSubmission)

	// User rating routes
	r.POST("/api/v1/users/:user_id/ratings", h.CreateAddUserRating)
	r.GET("/api/v1/users/:user_id/ratings", h.GetUserRatings)

	// Item category routes
	r.POST("/api/v1/item-categories", h.CreateItemCategory)

	// Statistics routes
	//r.GET("/api/v1/statistics", h.GetStatistics)

	// User activity monitoring routes
	r.GET("/api/v1/user-activity/:user_id", h.GetMonitoringUserActivity)

	// Eco-challenge routes
	r.POST("/api/v1/eco-challenges", h.CreateEcoChallenges)

	r.POST("/api/v1/eco-challenges/:challenge_id/participate", h.ParticipateChallenge)
	r.PUT("/api/v1/eco-challenges/:challenge_id/update-progress", h.UpdateEcoChallengeResult)

	// Eco-tips routes
	r.POST("/api/v1/eco-tips", h.CreateAddEcoTips)
	r.GET("/api/v1/eco-tips", h.GetAddEcoTips)

	return r
}
