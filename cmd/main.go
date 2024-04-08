package main

import (
	"effective/internal/app"
	"effective/pkg/logger"
)

func main() {
	if err := app.Start(); err != nil {
		logger.Error(err)
		return
	}
}
