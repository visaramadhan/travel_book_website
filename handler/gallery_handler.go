package handler

import (
	"net/http"
	"encoding/json"
	"github.com/golang-generic/service"
	"fmt"
)

type GalleryHandler struct {
	service service.GalleryService
}

func NewGalleryHandler(service service.GalleryService) *GalleryHandler {
	return &GalleryHandler{service: service}
}

// Handler function to return a list of photos, ensuring a minimum of 6
func (h *GalleryHandler) GetGalleryPhotos(w http.ResponseWriter, r *http.Request) {
	limit := 6 // set the minimum limit to 6 photos

	photos, err := h.service.GetGalleryPhotos(limit)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusBadRequest)
		return
	}

	// Respond with the photos only in JSON format (you can adjust how to display it based on your frontend)
	w.Header().Set("Content-Type", "application/json")

	// Preparing the response with only photos (you could convert photos to base64 or just serve them directly from a URL)
	var photoURLs []string
	for _, gallery := range photos {
		// Here we just assume the photo data is in a base64 string or URL, modify accordingly.
		// For base64 conversion, you can use encoding/base64 if needed.
		photoURLs = append(photoURLs, fmt.Sprintf("data:image/jpeg;base64,%s", gallery.Photo))
	}

	// Return the array of photo URLs in JSON format
	json.NewEncoder(w).Encode(photoURLs)
}