package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Drakoxw/go-cities-service/internal/models"
)

type MySQLCityRepository struct {
	DB *sql.DB
}

func NewMySQLCityRepository(db *sql.DB) *MySQLCityRepository {
	return &MySQLCityRepository{DB: db}
}

func (r *MySQLCityRepository) SearchCities(ctx context.Context, name string, page, limit int, sort, order string) ([]models.City, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT id, nombre, codigo_dane, departamento FROM cities WHERE nombre LIKE ? ORDER BY %s %s LIMIT ? OFFSET ?", sort, order)
	rows, err := r.DB.QueryContext(ctx, query, "%"+name+"%", limit, offset)
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
