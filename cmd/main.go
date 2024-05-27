package main

import (
	"log"

	"github.com/Drakoxw/go-cities-service/internal/cities/functions"
	"github.com/Drakoxw/go-cities-service/internal/cities/handlers"
	"github.com/Drakoxw/go-cities-service/internal/cities/utils"
	"github.com/Drakoxw/go-cities-service/internal/models"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var err error
	models.CitiesList, err = functions.LoadCitiesStart("data/cities.json")
	if err != nil {
		log.Fatalf("Error cargando datos: %v", err)
	}

	e.GET("/search", handlers.SearchCities)
	e.POST("/webhook/update-cities", handlers.UpdateCities)

	port := utils.GetPort()
	e.Start(port)
}
