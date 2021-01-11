package repository

import "github.com/mightyYaroslav/first-saga/order-service/internal/entity"

type CreateOrderParams struct {
	TicketId, Title, Description, Status string
	TotalPrice                           int
}

type Order interface {
	CreateOrder(params *CreateOrderParams) (*entity.Order, error)
	ApproveOrder(orderId string) error
	RejectOrder(orderId string) error
}
