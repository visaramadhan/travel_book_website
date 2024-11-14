package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-generic/database"
	"github.com/golang-generic/handler"
	"github.com/golang-generic/repository"
	"github.com/golang-generic/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Inisialisasi koneksi ke database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	} else {
		log.Println("Successfully connected to the database!")
	}
	defer db.Close()

	// Inisialisasi Place repository, service, dan handler
	placeRepository := repository.NewPlaceRepository(db)
	placeService := service.NewPlaceService(placeRepository)
	placeHandler := handler.NewPlaceHandler(placeService)

	// Inisialisasi Order repository, service, dan handler
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	// Inisialisasi Gallery
	galleryRepo := repository.NewGalleryRepository(db)
	galleryService := service.NewGalleryService(galleryRepo)
	galleryHandler := handler.NewGalleryHandler(galleryService)

	// Konfigurasi router menggunakan mux dengan prefix global "/api"
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter() // Membuat subrouter dengan prefix "/api"

	api.HandleFunc("/places", placeHandler.GetAllPlaces).Methods("GET")
	api.HandleFunc("/orders", orderHandler.HandleCreateOrder).Methods("POST")
	api.HandleFunc("/gallery", galleryHandler.GetGalleryPhotos).Methods("GET")
	// Menjalankan server pada port 8080
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
