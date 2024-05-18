package handlers

import (
	"net/http"
	"strconv"

	"github.com/Drakoxw/go-cities-service/internal/cities/usecase"
	"github.com/labstack/echo/v4"
)

type CityHandler struct {
	CityUC *usecase.CityUseCase
}

func NewCityHandler(e *echo.Echo, uc *usecase.CityUseCase) {
	handler := &CityHandler{CityUC: uc}
	e.GET("/search", handler.SearchCities)
}

func (h *CityHandler) SearchCities(c echo.Context) error {
	name := c.QueryParam("name")
	limitStr := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	cities, err := h.CityUC.SearchCitiesByName(name, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success":  true,
		"message":  "registros encontrados",
		"ciudades": cities,
	})
}
