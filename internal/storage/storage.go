package storage

import (
	"database/sql"
	"effective/config"
	"effective/pkg/logger"
	"fmt"
	"os"
	"path/filepath"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Name,
	)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	logger.Info(connString)
	err = db.Ping()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	direction := "up"

	err = ApplyMigrations(db, direction)
	if err != nil {
		return nil, fmt.Errorf("ApplyMigrations: %v", err)
	}
	return db, nil
}

func ApplyMigrations(db *sql.DB, direction string) error {
	migrationsDir := "migrations"

	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if direction == "up" && filepath.Ext(file.Name()) == ".up.sql" {
			err := applyMigration(db, filepath.Join(migrationsDir, file.Name()))
			if err != nil {
				return err
			}
		} else if direction == "down" && filepath.Ext(file.Name()) == ".down.sql" {
			err := applyMigration(db, filepath.Join(migrationsDir, file.Name()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func applyMigration(db *sql.DB, filePath string) error {
	migrationSQL, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		return err
	}

	logger.Info("Applied migration: %s\n", filePath)
	return nil
}
