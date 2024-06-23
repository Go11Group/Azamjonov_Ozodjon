package api

import (
	"client/api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes() *http.Server {
	mux := gin.Default()

	h := handler.NewHandler()

	// Stations
	mux.POST("/stations", h.Client)
	mux.GET("/stations/:id", h.Client)
	mux.PUT("/stations/:id", h.Client)
	mux.DELETE("/stations/:id", h.Client)
	mux.GET("/stations", h.Client)

	// Users
	mux.POST("/users", h.Client)
	mux.GET("/users/:id", h.Client)
	mux.PUT("/users/:id", h.Client)
	mux.DELETE("/users/:id", h.Client)
	mux.GET("/users", h.Client)

	// Cards
	mux.POST("/cards", h.Client)
	mux.GET("/cards/:id", h.Client)
	mux.PUT("/cards/:id", h.Client)
	mux.DELETE("/cards/:id", h.Client)
	mux.GET("/cards", h.Client)

	// Terminals
	mux.POST("/terminals", h.Client)
	mux.GET("/terminals/:id", h.Client)
	mux.PUT("/terminals/:id", h.Client)
	mux.DELETE("/terminals/:id", h.Client)
	mux.GET("/terminals", h.Client)

	// Transactions
	mux.POST("/transactions", h.Client)
	mux.GET("/transactions/:id", h.Client)
	mux.PUT("/transactions/:id", h.Client)
	mux.DELETE("/transactions/:id", h.Client)
	mux.GET("/transactions", h.Client)

	return &http.Server{Handler: mux, Addr: ":8081"}
}
