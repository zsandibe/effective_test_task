package app

import (
	"effective/config"
	"effective/internal/storage"
	"effective/pkg/logger"
	"fmt"
)

func Start() error {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		return fmt.Errorf("config.NewConfig: %v", err)
	}
	logger.Info("Config loaded successfully")

	db, err := storage.NewPostgresDB(cfg)
	if err != nil {
		return fmt.Errorf("config.NewConfig: %v", err)
	}

}
