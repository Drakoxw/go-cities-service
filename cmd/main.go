package main

import (
	"log"
	"time"

	"github.com/Drakoxw/go-cities-service/internal/cities/crones"
	"github.com/Drakoxw/go-cities-service/internal/cities/functions"
	"github.com/Drakoxw/go-cities-service/internal/cities/handlers"
	"github.com/Drakoxw/go-cities-service/internal/cities/utils"
	"github.com/Drakoxw/go-cities-service/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(utils.ForceJSONMiddleware)
	e.Use(middleware.Recover())

	var err error
	models.CitiesList, err = functions.LoadCitiesStart("data/cities.json")
	if err != nil {
		log.Fatalf("Error cargando datos: %v", err)
	}

	cron := time.NewTicker(12 * time.Hour)
	go func() {
		for {
			select {
			case <-cron.C:
				err = crones.UpdateCities()
				if err != nil {
					log.Println("❌❌" + err.Error() + "❌❌")
				}
			}
		}
	}()

	e.GET("/search", handlers.SearchCities)
	e.POST("/webhook/update-cities", handlers.UpdateCities)

	port := utils.GetPort()
	log.Fatal(e.Start(port))
}
