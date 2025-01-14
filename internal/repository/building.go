package repository

import (
	"building-service/internal/models"
	"database/sql"
	"strconv"
)

type BuildingRepository struct {
	db *sql.DB
}

func NewBuildingRepository(db *sql.DB) *BuildingRepository {
	return &BuildingRepository{db: db}
}

func (r *BuildingRepository) Create(building *models.Building) error {
	query := `
        INSERT INTO buildings (name, city, year_built, floors_count)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	return r.db.QueryRow(query,
		building.Name,
		building.City,
		building.YearBuilt,
		building.FloorsCount,
	).Scan(&building.ID)
}

func (r *BuildingRepository) List(city, year, floors string) ([]models.Building, error) {
	query := `SELECT id, name, city, year_built, floors_count FROM buildings WHERE 1=1`
	params := []interface{}{}

	if city != "" {
		query += " AND city = $" + strconv.Itoa(len(params)+1)
		params = append(params, city)
	}
	if year != "" {
		query += " AND year_built = $" + strconv.Itoa(len(params)+1)
		params = append(params, year)
	}
	if floors != "" {
		query += " AND floors_count = $" + strconv.Itoa(len(params)+1)
		params = append(params, floors)
	}

	rows, err := r.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var buildings []models.Building
	for rows.Next() {
		var b models.Building
		if err := rows.Scan(&b.ID, &b.Name, &b.City, &b.YearBuilt, &b.FloorsCount); err != nil {
			return nil, err
		}
		buildings = append(buildings, b)
	}
	return buildings, nil
}
