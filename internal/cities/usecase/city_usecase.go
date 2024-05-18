package usecase

import "github.com/Drakoxw/go-cities-service/internal/models"

type CityRepository interface {
	GetCitiesByName(name string, limit int) ([]models.City, error)
}

type CityUseCase struct {
	CityRepo CityRepository
}

func NewCityUseCase(repo CityRepository) *CityUseCase {
	return &CityUseCase{CityRepo: repo}
}

func (uc *CityUseCase) SearchCitiesByName(name string, limit int) ([]models.City, error) {
	return uc.CityRepo.GetCitiesByName(name, limit)
}
