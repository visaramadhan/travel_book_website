package service

import (
	"github.com/golang-generic/model"
	"github.com/golang-generic/repository"
)

// PlaceService is the service for working with places, tours, and galleries.
type PlaceService interface {
	GetPlaceWithTourAndGallery() ([]model.Place, error)
}

type placeService struct {
	repo repository.PlaceRepository
}

// NewPlaceService creates a new instance of PlaceService.
func NewPlaceService(repo repository.PlaceRepository) PlaceService {
	return &placeService{repo}
}

func (s *placeService) GetPlaceWithTourAndGallery() ([]model.Place, error) {
	return s.repo.GetPlaceWithTourAndGallery()
}
