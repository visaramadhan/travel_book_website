package repository

import (
	"database/sql"

	"github.com/golang-generic/model/dto"
)

// Inisialisasi interface OrderRepository
type OrderRepository interface {
    CreateOrder(order dto.OrderResponse) (int, error)
}

// Inisialisasi struct orderRepository
type orderRepository struct {
    db *sql.DB
}

// Fungsi NewOrderRepository untuk menginisialisasi orderRepository baru
func NewOrderRepository(db *sql.DB) OrderRepository {
    return &orderRepository{db}
}

// Implementasi metode CreateOrder
func (r *orderRepository) CreateOrder(order dto.OrderResponse) (int, error) {
    query := `INSERT INTO Orders (name, email, phone, date, number_of_ticket, message) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
    var orderID int

    err := r.db.QueryRow(query, order.Name, order.Email, order.PhoneNumber, order.OrderDate, order.NumberOfTicket, order.Message).Scan(&orderID)
    if err != nil {
        return 0, err
    }

    return orderID, nil
}
