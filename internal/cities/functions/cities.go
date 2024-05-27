package functions

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/Drakoxw/go-cities-service/internal/models"
)

func GetCitiesFromFile() ([]models.City, error) {
	// URL del archivo JSON
	url := "https://app.aveonline.co/assets/resources/public/listadociudades.json"

	// Obtener los datos desde la URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("No se pudieron obtener los datos desde la URL")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Respuesta no exitosa desde la URL")
	}

	var cities []models.City
	err = json.NewDecoder(resp.Body).Decode(&cities)
	if err != nil {
		return nil, errors.New("No se pudieron decodificar los datos")
	}

	// Escribir nuevos datos en el archivo JSON
	bytes, err := json.MarshalIndent(cities, "", "  ")
	if err != nil {
		return nil, errors.New("No se pudieron codificar los datos")
	}

	// Se reescribe el archivo tambn por si las dudas
	err = os.WriteFile("data/cities.json", bytes, 0644)
	if err != nil {
		return nil, errors.New("No se pudieron guardar los datos")
	}

	return cities, nil

}
