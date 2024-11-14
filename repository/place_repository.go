package repository

import (
	"database/sql"
	"fmt"

	"github.com/golang-generic/model"
)

// PlaceRepository is the interface for working with places, tours, and galleries.
type PlaceRepository interface {
	GetPlaceWithTourAndGallery() ([]model.Place, error)
}

// placeRepository is a concrete implementation of the PlaceRepository interface.
type placeRepository struct {
	db *sql.DB
}

// NewPlaceRepository creates a new instance of PlaceRepository.
func NewPlaceRepository(db *sql.DB) PlaceRepository {
	return &placeRepository{db}
}

// GetPlaceWithTourAndGallery retrieves data about places, tours, and gallery photos.
func (r *placeRepository) GetPlaceWithTourAndGallery() ([]model.Place, error) {
	query := `
		SELECT
    p.id AS place_id,
    p.name AS place_name,
    p.description AS place_description,
    p.photo AS place_thumbnail,
    p.price AS place_price,
    t.name AS tour_name,
    t.date AS tour_date,
    g.name AS gallery_name,
    g.photo AS gallery_photo
FROM
    place p
JOIN
    tour t ON p.id = t.place_id
JOIN (
    SELECT 
        g.id AS id,
        g.name AS name,
        g.photo AS photo,
        g.id_place,
        ROW_NUMBER() OVER (PARTITION BY g.id_place ORDER BY g.id) AS rn
    FROM gallery g
) g ON p.id = g.id_place
WHERE g.rn <= 6
ORDER BY t.date DESC;

	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	var places []model.Place
	for rows.Next() {
		var place model.Place
		err := rows.Scan(
			&place.ID,
			&place.Name,
			&place.Description,
			&place.Price,
			&place.Photo,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		places = append(places, place)
	}

	return places, nil
}
