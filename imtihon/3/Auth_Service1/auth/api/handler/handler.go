package handler

import (
	"auth_service/service"
	"log/slog"
)

type Handler struct {
	Logger *slog.Logger
	User   *service.UserService
}

func NewHandler(logger *slog.Logger, user *service.UserService) *Handler {
	return &Handler{Logger: logger, User: user}
}
