package handler

import (
    "encoding/json"
    "net/http"

    "github.com/golang-generic/model/dto"
    "github.com/golang-generic/service"
)

// OrderHandler struct untuk menangani permintaan HTTP terkait pesanan
type OrderHandler struct {
    orderService service.OrderService
}

// NewOrderHandler membuat instance baru dari OrderHandler
func NewOrderHandler(orderService service.OrderService) *OrderHandler {
    return &OrderHandler{orderService}
}

// HandleCreateOrder menangani permintaan POST untuk membuat pesanan baru
func (h *OrderHandler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
    var order dto.OrderResponse

    // Decode JSON request body ke struct Order
    if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Memanggil service untuk membuat pesanan
    orderID, err := h.orderService.CreateOrder(order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Mengembalikan respon sukses dengan ID pesanan baru
    response := map[string]interface{}{
        "message": "Order created successfully",
        "order_id": orderID,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
