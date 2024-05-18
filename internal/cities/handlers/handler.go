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
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	sort := c.QueryParam("sort")
	order := c.QueryParam("order")

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	if sort == "" {
		sort = "nombre"
	}
	if order == "" {
		order = "ASC"
	}

	cities, err := h.CityUC.SearchCities(c.Request().Context(), name, page, limit, sort, order)
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
