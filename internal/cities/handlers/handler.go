package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Drakoxw/go-cities-service/internal/cities/functions"
	"github.com/Drakoxw/go-cities-service/internal/cities/usecase"
	"github.com/Drakoxw/go-cities-service/internal/cities/utils"
	"github.com/Drakoxw/go-cities-service/internal/models"
	"github.com/labstack/echo/v4"
)

type CityHandler struct {
	CityUC *usecase.CityUseCase
}

func NewCityHandler(e *echo.Echo, uc *usecase.CityUseCase) {
	handler := &CityHandler{CityUC: uc}
	e.GET("/search", handler.SearchCities)
	e.POST("/webhook/update-cities", handler.UpdateCities)
}

func (h *CityHandler) UpdateCities(c echo.Context) error {

	cities, err := functions.GetCitiesFromFile()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BabResponse(err.Error()))
	}

	if err := h.CityUC.UpdateCities(c.Request().Context(), cities); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BabResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utils.OkResponse("Ciudades actualizadas correctamente"))
}

func (h *CityHandler) SearchCities(c echo.Context) error {
	name := c.QueryParam("query")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	sort := c.QueryParam("sort")
	order := c.QueryParam("order")

	if len(name) < 4 {
		return c.JSON(http.StatusBadRequest, utils.BabResponse("Se requie al menos 3 caracteres"))
	}

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

	validSorts := map[string]bool{"id": true, "nombre": true, "codigodane": true, "departamento": true}
	if !validSorts[sort] {
		return c.JSON(http.StatusBadRequest, utils.BabResponse("Parámetro de ordenación inválido"))
	}

	order = strings.ToUpper(order)
	if order != "ASC" && order != "DESC" {
		return c.JSON(http.StatusBadRequest, utils.BabResponse("Parámetro de orden inválido"))
	}

	cities, err := h.CityUC.SearchCities(c.Request().Context(), name, page, limit, sort, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BabResponse(err.Error()))
	}

	message := "registros encontrados"
	if cities == nil {
		message = "no hubieron coincidencias"
		cities = []models.City{}
	}

	return c.JSON(http.StatusOK, utils.OkResponseData(message, cities))
}
