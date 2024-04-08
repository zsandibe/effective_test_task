package service

import (
	"context"
	"effective/internal/domain"
)

func (s *serviceCar) AddCar(ctx context.Context, car domain.Car) error {
	checkedCar, err := getCarInfoByRegNum(car.RegNum)
	if err != nil {
		return err
	}
	return s.repo.AddCar(ctx, checkedCar)
}

func (s *serviceCar) GetCarById(ctx context.Context, id int) (domain.Car, error) {
	return s.repo.GetCarById(ctx, id)
}

func (s *serviceCar) GetCarsList(ctx context.Context, params domain.CarsListParams) ([]domain.Car, error) {
	return s.repo.GetCarsList(ctx, params)
}

func (s *serviceCar) UpdateCarInfo(ctx context.Context, carID int, params domain.CarDataUpdatingRequest) error {
	return s.repo.UpdateCarInfo(ctx, carID, params)
}
