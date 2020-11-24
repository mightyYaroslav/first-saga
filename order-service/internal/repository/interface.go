package repository

import "github.com/mightyYaroslav/first-saga/order-service/internal/entity"

type Order interface {
	CreateOrder(ticketId, title, description, status string, totalPrice int) (*entity.Order, error)
	ApproveOrder(orderId string) error
	RejectOrder(orderId string) error
}
