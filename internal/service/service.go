package service

import (
	"context"
	"effective/internal/domain"
	"effective/internal/repository"
)

type Service interface {
	AddCar(ctx context.Context, car domain.Car) error
	GetCarById(ctx context.Context, id int) (domain.Car, error)
	GetCarsList(ctx context.Context, params domain.CarsListParams) ([]domain.Car, error)
	UpdateCarInfo(ctx context.Context, carID int, params domain.CarDataUpdatingRequest) error
}

type serviceCar struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *serviceCar {
	return &serviceCar{
		repo: repo,
	}
}
