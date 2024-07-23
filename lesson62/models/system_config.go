package models

import (
	"database/sql"

	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type SystemConfig struct {
	Logger   *zap.Logger
	Enforcer *casbin.Enforcer
	DBR      *redis.Client
	DBP      *sql.DB
}
