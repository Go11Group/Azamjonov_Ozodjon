package main

import (
	"fmt"
	"lesson62/api"
	"lesson62/pkg/systemconfig"
	"log"

	"go.uber.org/zap"
)

func main() {
	systemConfig, err := systemconfig.NewSystemConfig()
	if err != nil {
		log.Panicf("Error with creating system config: %s", err)
		return
	}

	router := api.NewRouter(systemConfig)

	if err := router.Run(":5555"); err != nil {
		systemConfig.Logger.Panic("Error with running router", zap.Error(err))
	}
	fmt.Println("server is running on port 5555...")
}
