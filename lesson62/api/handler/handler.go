package handler

import (
	"lesson62/models"
	rd "lesson62/storage/redis"

	"go.uber.org/zap"
)

type Handler struct {
	Restaurant *rd.RestaurantRepo
	Users      *rd.UsersRepo
	Logger     *zap.Logger
}

func NewHandler(sysconf *models.SystemConfig) *Handler {
	users := rd.NewUsersClient(sysconf.DBR)
	restaurant := rd.NewRestaurantRepo(sysconf.DBR)

	return &Handler{
		Restaurant: restaurant,
		Users:      users,
		Logger:     sysconf.Logger,
	}
}
