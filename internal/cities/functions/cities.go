package functions

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Drakoxw/go-cities-service/internal/models"
)

func GetCitiesFromFile() ([]models.City, error) {
	url := "https://app.aveonline.co/assets/resources/public/listadociudades.json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Error al descargar el archivo JSON")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error al leer el archivo JSON")
	}

	var cities []models.City
	if err := json.Unmarshal(body, &cities); err != nil {
		return nil, errors.New("Error al decodificar el archivo JSON")
	}

	return cities, nil
}
