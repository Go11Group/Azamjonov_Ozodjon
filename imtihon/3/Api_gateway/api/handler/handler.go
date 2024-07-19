package handler

import (
	item "github.com/imtihon/3/Api_gateway/generated/generated/items_service"
	auth "github.com/imtihon/3/Api_gateway/generated/genproto"
	"log/slog"
)

type Handler struct {
	Logger *slog.Logger
	Auth   auth.AuthServiceClient
	Items  item.ItemsServiceClient
}

func NewHandler(itemClient item.ItemsServiceClient, authClient auth.AuthServiceClient, logger *slog.Logger) *Handler {
	return &Handler{
		Auth:   authClient,
		Logger: logger,
		Items:  itemClient,
	}
}
