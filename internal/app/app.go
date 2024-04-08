package app

import (
	"context"
	"effective/config"
	v1 "effective/internal/delivery/http/v1"
	"effective/internal/delivery/server"
	"effective/internal/repository"
	"effective/internal/service"
	"effective/internal/storage"
	"effective/pkg/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	repository := repository.NewRepository(db)
	logger.Info("Repository loaded successfully")

	service := service.NewService(repository)
	logger.Info("Service loaded successfully")

	delivery := v1.NewHandler(service)
	logger.Info("Delivery loaded successfully")

	server := server.NewServer(cfg, delivery.Routes())
	go func() {
		if err := server.Run(); err != nil {
			logger.Error("failed to start server: %v", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	select {
	case <-quit:
		logger.Info("Received interrupt signal. Shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Error("Error during server shutdown: ", err)
		}

		logger.Info("Server gracefully stopped")
	}
	return nil
}
