package model

// Gallery struct untuk menyimpan informasi tentang gambar galeri
type Gallery struct {
	ID    int    `json:"id"`    // ID foto 
	Name	string  `json:"name"`
	Photo []byte `json:"photo"` // Foto dalam bentuk byte array
}
