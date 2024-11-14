package repository

import (
	"database/sql"
	"fmt"
	"github.com/golang-generic/model"
)

type GalleryRepository interface {
	GetGalleryPhotos(limit int) ([]model.Gallery, error)
}

type galleryRepository struct {
	db *sql.DB
}

func NewGalleryRepository(db *sql.DB) GalleryRepository {
	return &galleryRepository{db}
}

// Method to retrieve photos from the Gallery table ensuring at least 6 photos
func (r *galleryRepository) GetGalleryPhotos(limit int) ([]model.Gallery, error) {
	var galleries []model.Gallery
	query := "SELECT photo FROM Gallery LIMIT $1"

	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var gallery model.Gallery
		err := rows.Scan(&gallery.Photo)
		if err != nil {
			return nil, err
		}
		galleries = append(galleries, gallery)
	}

	// Validate that at least 6 photos are returned
	if len(galleries) < 6 {
		return nil, fmt.Errorf("insufficient photos: at least 6 photos are required")
	}

	return galleries, nil
}