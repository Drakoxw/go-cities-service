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
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dsn := "root:@tcp(localhost:3307)/drakodb" // conexion local
	// dsn := "root:@tcp(mysql:3306)/drakodb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.CreateTable(db)
	database.LoadData(db)

	e := echo.New()
	e.Use(middleware.Recover())

	cityRepo := mysql.NewMySQLCityRepository(db)
	cityUC := usecase.NewCityUseCase(cityRepo)
	handlers.NewCityHandler(e, cityUC)

	cron := time.NewTicker(30 * time.Second)
	go func() {
		for {
			select {
			case <-cron.C:
				err = crones.UpdateCities(cityUC)
				if err != nil {
					// log.Println("❌❌" + err.Error() + "❌❌")
				}
				log.Println("Datos actualizados")
			}
		}
	}()

	port := utils.GetPort()
	log.Fatal(e.Start(port))
}
