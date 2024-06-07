package handler

import (
	"net/http"
	"users_and_produncts/storage/postgres"
)

type Handler struct {
	Users         *postgres.Users
	Products      *postgres.Products
}

func NewHandler(users *postgres.Users, products *postgres.Products) *Handler {
	return &Handler{
		Users:         users,
		Products:      products,
	}
}

func CreateServer(handler *Handler) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /GetUser", handler.GetUsers)
	mux.HandleFunc("POST /CreateUser", handler.CreateUser)
	mux.HandleFunc("PUT /UpdateUser", handler.UpdateUser)
	mux.HandleFunc("DELETE /DeleteUser", handler.DeleteUser)

	mux.HandleFunc("GET /GetProduct", handler.GetProduct)
	mux.HandleFunc("POST /CreateProduct", handler.CreateProduct)
	mux.HandleFunc("PUT /UpdateProduct", handler.UpdateProduct)
	mux.HandleFunc("DELETE /DeleteProduct", handler.DeleteProduct)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
