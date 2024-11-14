package service

import (
	"github.com/golang-generic/repository"
	"github.com/golang-generic/model"
)

type GalleryService interface {
	GetGalleryPhotos(limit int) ([]model.Gallery, error)
}

type galleryService struct {
	repo repository.GalleryRepository
}

func NewGalleryService(repo repository.GalleryRepository) GalleryService {
	return &galleryService{repo}
}

// Service method to get photos from the repository
func (s *galleryService) GetGalleryPhotos(limit int) ([]model.Gallery, error) {
	return s.repo.GetGalleryPhotos(limit)
}