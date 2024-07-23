package systemconfig

import (
	"fmt"
	"lesson62/models"
	"lesson62/pkg/logger"
	"lesson62/storage/postgres"
	"lesson62/storage/redis"

	pgadapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
)

func NewSystemConfig() (*models.SystemConfig, error) {
	// create Logger
	logger, err := logger.New("debug", "app.log")
	if err != nil {
		return nil, err
	}

	// create redis database
	dbr, err := redis.ConnectDB()
	if err != nil {
		return nil, err
	}

	// create postgres database
	dbp, err := postgres.ConnectDB()
	if err != nil {
		return nil, err
	}

	// create new casbin enforcer
	a, err := pgadapter.NewAdapter(dbp, "postgres", "casbin_rule")
	if err != nil {
		return nil, fmt.Errorf("error with creating adapter: %s", err)
	}

	e, err := casbin.NewEnforcer("/home/jons/go/src/github.com/Azamjonov_Ozodjon/lesson62/config/model.conf", a)
	if err != nil {
		return nil, fmt.Errorf("error with creating enforcer: %s", err)
	}

	err = e.LoadPolicy()
	if err != nil {
		return nil, fmt.Errorf("error with loading enforcer policy: %s", err)
	}

	return &models.SystemConfig{
		Logger:   logger,
		DBR:      dbr,
		DBP:      dbp,
		Enforcer: e,
	}, nil
}
