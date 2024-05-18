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

func (r *MySQLCityRepository) UpdateCities(ctx context.Context, cities []models.City) error {
	if len(cities) == 0 {
		return fmt.Errorf("se quiere al menos una ciudad")
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "DELETE FROM cities"); err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO cities (nombre, codigo_dane, departamento) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	batchSize := 500
	for i := 0; i < len(cities); i += batchSize {
		end := i + batchSize
		if end > len(cities) {
			end = len(cities)
		}

		batch := cities[i:end]
		for _, city := range batch {
			if _, err := stmt.ExecContext(ctx, city.Nombre, city.CodigoDANE, city.Departamento); err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}
