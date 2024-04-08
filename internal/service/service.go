package service

import "effective/internal/repository"

type Service interface {
}

type serviceCar struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *serviceCar {
	return &serviceCar{
		repo: repo,
	}
}
