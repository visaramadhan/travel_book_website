package dto

import "time"

type OrderResponse struct {
	OrderID        int       `json:"id"`
	Name           string    `json:"nama"`
	Email          string    `json:"email"`
	ConfirmEmail	string	`json:"confirmEmail"`
	PhoneNumber    int       `json:"phone"`
	NumberOfTicket int       `json:"number_of_ticket"`
	OrderDate      time.Time `json:"date"`
	TourID         int       `json:"id_tour"`
	Message			string	`json:"message"`
}
