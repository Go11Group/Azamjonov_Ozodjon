package api

import (
	"lesson62/api/handler"
	"lesson62/api/middleware"
	"lesson62/models"

	"github.com/gin-gonic/gin"
)

func NewRouter(sysconf *models.SystemConfig) *gin.Engine {
	r := gin.Default()

	h := handler.NewHandler(sysconf)
	r.Use(middleware.CheckPermissionMiddleware(sysconf.Enforcer))

	restaurant := r.Group("/restaurant")
	restaurant.POST("/create", h.CreateRestaurant)
	restaurant.GET("/getall", h.GetAllRestaurants)
	restaurant.GET("/:id", h.GetById)
	restaurant.PUT("/update", h.UpdateRestaurant)
	restaurant.DELETE("/:id", h.DeleteRestaurant)

	users := r.Group("/users")
	users.POST("/create", h.CreateUser)
	users.GET("/getall", h.GetAllUsers)
	users.GET("/:id", h.GetUserById)
	users.PUT("/update", h.UpdateUser)
	users.DELETE("/:id", h.DeleteUser)

	return r
}
