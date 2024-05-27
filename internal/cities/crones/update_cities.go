package crones

import (
	"github.com/Drakoxw/go-cities-service/internal/cities/functions"
	"github.com/Drakoxw/go-cities-service/internal/models"
)

func UpdateCities() {
	var err error
	models.CitiesList, err = functions.LoadCitiesStart("data/cities.json")
	if err != nil {
		println(err.Error())
	}
}
