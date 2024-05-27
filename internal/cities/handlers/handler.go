package handlers

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/Drakoxw/go-cities-service/internal/cities/functions"
	"github.com/Drakoxw/go-cities-service/internal/cities/utils"
	"github.com/Drakoxw/go-cities-service/internal/models"
	"github.com/labstack/echo/v4"
)

func UpdateCities(c echo.Context) error {
	cities, err := functions.GetCitiesFromFile()
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.BabResponse(err.Error()))
	}

	models.CitiesList = cities

	return c.JSON(http.StatusOK, utils.OkResponse("Datos actualizados"))
}

func SearchCities(c echo.Context) error {
	query := c.QueryParam("query")
	limitParam := c.QueryParam("limit")
	pageParam := c.QueryParam("page")
	sortParam := c.QueryParam("sort")
	orderParam := c.QueryParam("order")

	if len(query) < 4 {
		return c.JSON(http.StatusBadRequest, utils.BabResponse("Se requiere al menos 3 caracteres"))
	}

	// Validar los sort permitidos
	validSorts := map[string]bool{"nombre": true, "codigodane": true, "departamento": true}
	if !validSorts[sortParam] {
		return c.JSON(http.StatusBadRequest, utils.BabResponse("Parámetro de busqueda inválido"))
	}

	orderParam = strings.ToUpper(orderParam)
	if orderParam != "ASC" && orderParam != "DESC" {
		return c.JSON(http.StatusBadRequest, utils.BabResponse("Parámetro de orden inválido"))
	}

	var resultado []models.City

	// Filtrar ciudades por el query
	for _, ciudad := range models.CitiesList {
		dataFilteredForKey := ciudad.Nombre
		switch sortParam {
		case "codigodane":
			dataFilteredForKey = ciudad.CodigoDANE
		case "departamento":
			dataFilteredForKey = ciudad.Departamento
		}
		if strings.Contains(strings.ToLower(dataFilteredForKey), strings.ToLower(query)) {
			resultado = append(resultado, ciudad)
		}
	}

	// Ordenar resultados
	if sortParam != "" {
		switch sortParam {
		case "codigodane":
			sort.Slice(resultado, func(i, j int) bool {
				if orderParam == "DESC" {
					return resultado[i].CodigoDANE > resultado[j].CodigoDANE
				}
				return resultado[i].CodigoDANE < resultado[j].CodigoDANE
			})
		case "nombre":
			sort.Slice(resultado, func(i, j int) bool {
				if orderParam == "DESC" {
					return resultado[i].Nombre > resultado[j].Nombre
				}
				return resultado[i].Nombre < resultado[j].Nombre
			})
		case "departamento":
			sort.Slice(resultado, func(i, j int) bool {
				if orderParam == "DESC" {
					return resultado[i].Departamento > resultado[j].Departamento
				}
				return resultado[i].Departamento < resultado[j].Departamento
			})
		}
	}

	// Paginación y límite
	limit := 10
	if limitParam != "" {
		parsedLimit, err := strconv.Atoi(limitParam)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	page := 1
	if pageParam != "" {
		parsedPage, err := strconv.Atoi(pageParam)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(resultado) {
		start = len(resultado)
	}
	if end > len(resultado) {
		end = len(resultado)
	}

	return c.JSON(http.StatusOK, utils.OkResponseData("Data found", resultado[start:end]))
}
