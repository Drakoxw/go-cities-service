package mysql

import (
	"database/sql"

	"github.com/Drakoxw/go-cities-service/internal/models"
)

type MySQLCityRepository struct {
	DB *sql.DB
}

func NewMySQLCityRepository(db *sql.DB) *MySQLCityRepository {
	return &MySQLCityRepository{DB: db}
}

func (r *MySQLCityRepository) GetCitiesByName(name string, limit int) ([]models.City, error) {
	rows, err := r.DB.Query("SELECT id, nombre, codigo_dane, departamento FROM cities WHERE nombre LIKE ? LIMIT ?", "%"+name+"%", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []models.City
	for rows.Next() {
		var city models.City
		if err := rows.Scan(&city.ID, &city.Nombre, &city.CodigoDANE, &city.Departamento); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	return cities, nil
}
