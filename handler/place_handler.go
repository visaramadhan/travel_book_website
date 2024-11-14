package handler

import (
	"encoding/json"
	"net/http"

	"github.com/golang-generic/service"
)

// PlaceHandler is the HTTP handler for places, tours, and galleries.
type PlaceHandler struct {
	service service.PlaceService
}

// NewPlaceHandler creates a new instance of PlaceHandler.
func NewPlaceHandler(service service.PlaceService) *PlaceHandler {
	return &PlaceHandler{service}
}

// GetAllPlaces is the handler function to fetch places with tours and gallery.
func (h *PlaceHandler) GetAllPlaces(w http.ResponseWriter, r *http.Request) {
	places, err := h.service.GetPlaceWithTourAndGallery()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(places); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
