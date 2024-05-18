package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Drakoxw/go-cities-service/internal/cities/crones"
	"github.com/Drakoxw/go-cities-service/internal/cities/handlers"
	"github.com/Drakoxw/go-cities-service/internal/cities/repository/mysql"
	"github.com/Drakoxw/go-cities-service/internal/cities/usecase"
	"github.com/Drakoxw/go-cities-service/internal/cities/utils"
	"github.com/Drakoxw/go-cities-service/internal/database"
	"github.com/labstack/echo/v4"
)

func main() {
	// dsn := "root:@tcp(localhost:3307)/drakodb" // conexion local
	dsn := "root:@tcp(mysql:3306)/drakodb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.CreateTable(db)
	database.LoadData(db)

	e := echo.New()

	cityRepo := mysql.NewMySQLCityRepository(db)
	cityUC := usecase.NewCityUseCase(cityRepo)
	handlers.NewCityHandler(e, cityUC)

	port := utils.GetPort()
	go func() {
		log.Fatal(e.Start(port))
	}()

	cron := time.NewTicker(12 * time.Hour)
	go func() {
		for {
			select {
			case <-cron.C:
				err = crones.UpdateCities(cityUC)
				if err != nil {
					log.Println("❌❌" + err.Error() + "❌❌")
				}
			}
		}
	}()

	select {}
}
