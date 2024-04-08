package main

import (
	"effective/internal/app"
	"effective/pkg/logger"
)

// @title EffectiveTask API
// @description This is basic server for a car service
// @version 1.0
// @host localhost:8888
// @BasePath /api/v1
func main() {
	if err := app.Start(); err != nil {
		logger.Error(err)
		return
	}
}
