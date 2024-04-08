package repository

import (
	"context"
	"database/sql"
	"effective/internal/domain"
)

type Repository interface {
	AddCar(ctx context.Context, car domain.Car) error
	GetCarById(ctx context.Context, id int) (domain.Car, error)
	GetCarsList(ctx context.Context, params domain.CarsListParams) ([]domain.Car, error)
	UpdateCarInfo(ctx context.Context, carID int, params domain.CarDataUpdatingRequest) error
	DeleteCarById(ctx context.Context, id int) error
}

type repositoryPostgres struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repositoryPostgres {
	return &repositoryPostgres{
		db: db,
	}
}
