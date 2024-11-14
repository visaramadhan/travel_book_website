package service

import (
	"errors"
	"time"

	"github.com/golang-generic/model/dto"
	"github.com/golang-generic/repository"
)

// OrderService interface untuk layanan pesanan
type OrderService interface {
    CreateOrder(order dto.OrderResponse) (int, error)
}

// struct orderService
type orderService struct {
    orderRepo repository.OrderRepository
}

// NewOrderService untuk membuat instance OrderService baru
func NewOrderService(orderRepo repository.OrderRepository) OrderService {
    return &orderService{orderRepo}
}

// Implementasi method CreateOrder di orderService
func (s *orderService) CreateOrder(order dto.OrderResponse) (int, error) {
    // Validasi email dan konfirmasi email
    if order.Email != order.ConfirmEmail {
        return 0, errors.New("email dan confirm email harus sama")
    }

    // Validasi jumlah tiket
    if order.NumberOfTicket <= 0 {
        return 0, errors.New("jumlah tiket harus lebih dari 0")
    }

    // Set tanggal pesanan ke tanggal saat ini jika belum diisi
    if order.OrderDate.IsZero() {
        order.OrderDate = time.Now()
    }

    // Memanggil repository untuk menyimpan pesanan
    return s.orderRepo.CreateOrder(order)
}