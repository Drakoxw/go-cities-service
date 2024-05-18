package usecase

import (
	"context"

	"github.com/Drakoxw/go-cities-service/internal/models"
)

type CityRepository interface {
	SearchCities(ctx context.Context, name string, page, limit int, sort, order string) ([]models.City, error)
	UpdateCities(ctx context.Context, cities []models.City) error
}

type CityUseCase struct {
	CityRepo CityRepository
}

func NewCityUseCase(repo CityRepository) *CityUseCase {
	return &CityUseCase{CityRepo: repo}
}

func (u *CityUseCase) UpdateCities(ctx context.Context, cities []models.City) error {
	return u.CityRepo.UpdateCities(ctx, cities)
}

func (u *CityUseCase) SearchCities(ctx context.Context, name string, page, limit int, sort, order string) ([]models.City, error) {
	return u.CityRepo.SearchCities(ctx, name, page, limit, sort, order)
}
