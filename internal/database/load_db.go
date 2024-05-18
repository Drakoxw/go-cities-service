package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Drakoxw/go-cities-service/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

func CreateTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS cities (
        id INT AUTO_INCREMENT PRIMARY KEY,
        nombre VARCHAR(255) NOT NULL,
        codigo_dane VARCHAR(20) NOT NULL,
        departamento VARCHAR(255) NOT NULL
    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadData(db *sql.DB) {
	// Leer el archivo JSON
	jsonFile, err := os.Open("data/cities.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var cities []models.City
	json.Unmarshal(byteValue, &cities)

	// Insertar datos en la base de datos
	for _, city := range cities {
		_, err = db.Exec("INSERT INTO cities (nombre, codigo_dane, departamento) VALUES (?, ?, ?)",
			city.Nombre, city.CodigoDANE, city.Departamento)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Datos cargados exitosamente en la base de datos")
}
