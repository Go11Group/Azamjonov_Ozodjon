package api

import (
	"atto/api/handler"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *http.Server {
	mux := gin.Default()

	h := handler.NewHandler(db)

	mux.POST("/stations", h.CreateStation)
	mux.GET("/stations/:id", h.GetStation)
	mux.PUT("/stations/:id", h.UpdateStation)
	mux.DELETE("/stations/:id", h.DeleteStation)
	mux.GET("/stations", h.GetAllStations)

	mux.POST("/users", h.CreateUser)
	mux.GET("/users/:id", h.GetUser)
	mux.PUT("/users/:id", h.UpdateUser)
	mux.DELETE("/users/:id", h.DeleteUser)
	mux.GET("/users", h.GetAllUsers)

	mux.POST("/cards", h.CreateCard)
	mux.GET("/cards/:id", h.GetCard)
	mux.PUT("/cards/:id", h.UpdateCard)
	mux.DELETE("/cards/:id", h.DeleteCard)
	mux.GET("/cards", h.GetAllCards)

	mux.POST("/terminals", h.CreateTerminal)
	mux.GET("/terminals/:id", h.GetTerminal)
	mux.PUT("/terminals/:id", h.UpdateTerminal)
	mux.DELETE("/terminals/:id", h.DeleteTerminal)
	mux.GET("/terminals", h.GetAllTerminals)

	mux.POST("/transactions", h.CreateTransaction)
	mux.GET("/transactions/:id", h.GetTransaction)
	mux.PUT("/transactions/:id", h.UpdateTransaction)
	mux.DELETE("/transactions/:id", h.DeleteTransaction)
	mux.GET("/transactions", h.GetAllTransactions)

	return &http.Server{Handler: mux, Addr: ":8080"}
}
