package crones

import (
	"context"

	"github.com/Drakoxw/go-cities-service/internal/cities/functions"
	"github.com/Drakoxw/go-cities-service/internal/cities/handlers"
	"github.com/Drakoxw/go-cities-service/internal/cities/usecase"
)

func UpdateCities(uc *usecase.CityUseCase) error {

	h := &handlers.CityHandler{CityUC: uc}

	ctx := context.Background()
	cities, err := functions.GetCitiesFromFile()
	if err != nil {
		return err
	}

	if err := h.CityUC.UpdateCities(ctx, cities); err != nil {
		return err
	}

	return nil
}
