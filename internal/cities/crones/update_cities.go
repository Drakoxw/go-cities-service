package crones

import (
	"github.com/Drakoxw/go-cities-service/internal/cities/functions"
	"github.com/Drakoxw/go-cities-service/internal/models"
)

func UpdateCities() error {
	cities, err := functions.GetCitiesFromFile()
	if err != nil {
		return err
	}
	models.CitiesList = cities
	return nil
}
