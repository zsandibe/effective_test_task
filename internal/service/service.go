package service

import (
	"context"
	"effective/config"
	"effective/internal/domain"
	"effective/internal/repository"
)

type Service interface {
	AddCar(ctx context.Context, car domain.Car) error
	GetCarById(ctx context.Context, id int) (domain.Car, error)
	GetCarsList(ctx context.Context, params domain.CarsListParams) ([]domain.Car, error)
	UpdateCarInfo(ctx context.Context, carID int, params domain.CarDataUpdatingRequest) error
	DeleteCarById(ctx context.Context, id int) error
}

type serviceCar struct {
	repo repository.Repository
	conf config.Config
}

func NewService(repo repository.Repository) *serviceCar {
	return &serviceCar{
		repo: repo,
	}
}
