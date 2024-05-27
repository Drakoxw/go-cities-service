package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Drakoxw/go-cities-service/internal/cities/utils"
	"github.com/labstack/echo/v4"
)

type Ciudad struct {
	CodigoDANE   string `json:"codigodane"`
	Nombre       string `json:"nombre"`
	Departamento string `json:"departamento"`
}

func cargarDatos(ruta string) ([]Ciudad, error) {
	var ciudades []Ciudad

	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	bytes, err := io.ReadAll(archivo)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &ciudades)
	if err != nil {
		return nil, err
	}

	return ciudades, nil
}

func buscarCiudades(c echo.Context) error {
	query := c.QueryParam("query")
	limitParam := c.QueryParam("limit")
	pageParam := c.QueryParam("page")
	sortParam := c.QueryParam("sort")
	orderParam := c.QueryParam("order")

	if len(query) < 4 {
		return c.JSON(http.StatusBadRequest, utils.BabResponse("Se requiere al menos 3 caracteres"))
	}

	var resultado []Ciudad

	// Filtrar ciudades por el query
	for _, ciudad := range ciudades {
		if query == "" || strings.Contains(strings.ToLower(ciudad.Nombre), strings.ToLower(query)) {
			resultado = append(resultado, ciudad)
		}
	}

	// Ordenar resultados
	if sortParam != "" {
		// Validar los sort permitidos
		validSorts := map[string]bool{"nombre": true, "codigodane": true, "departamento": true}
		if !validSorts[sortParam] {
			return c.JSON(http.StatusBadRequest, utils.BabResponse("Parámetro de busqueda inválido"))
		}

		switch sortParam {
		case "codigodane":
			sort.Slice(resultado, func(i, j int) bool {
				if orderParam == "desc" {
					return resultado[i].CodigoDANE > resultado[j].CodigoDANE
				}
				return resultado[i].CodigoDANE < resultado[j].CodigoDANE
			})
		case "nombre":
			sort.Slice(resultado, func(i, j int) bool {
				if orderParam == "desc" {
					return resultado[i].Nombre > resultado[j].Nombre
				}
				return resultado[i].Nombre < resultado[j].Nombre
			})
		case "departamento":
			sort.Slice(resultado, func(i, j int) bool {
				if orderParam == "desc" {
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

var ciudades []Ciudad

func actualizarCiudades(c echo.Context) error {
	// URL del archivo JSON
	url := "https://app.aveonline.co/assets/resources/public/listadociudades.json"

	// Obtener los datos desde la URL
	resp, err := http.Get(url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BabResponse("No se pudieron obtener los datos desde la URL"))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, utils.BabResponse("Respuesta no exitosa desde la URL"))
	}

	var nuevasCiudades []Ciudad
	err = json.NewDecoder(resp.Body).Decode(&nuevasCiudades)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BabResponse("No se pudieron decodificar los datos"))
	}

	// Actualizar datos en memoria
	ciudades = nuevasCiudades

	// Escribir nuevos datos en el archivo JSON
	bytes, err := json.MarshalIndent(nuevasCiudades, "", "  ")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BabResponse("No se pudieron codificar los datos"))
	}

	// Se reescribe el archivo tambn por si las dudas
	err = os.WriteFile("data/cities.json", bytes, 0644)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BabResponse("No se pudieron guardar los datos"))
	}

	return c.JSON(http.StatusOK, utils.OkResponse("Datos actualizados"))
}

func main() {
	e := echo.New()

	_, err := cargarDatos("data/cities.json")
	if err != nil {
		log.Fatalf("Error cargando datos: %v", err)
	}

	e.GET("/search", buscarCiudades)
	e.POST("/webhook/update-cities", actualizarCiudades)

	e.Start(":8989")
}
