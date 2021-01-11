package usecase

import (
	"github.com/mightyYaroslav/first-saga/order-service/internal/entity"
	"github.com/mightyYaroslav/first-saga/order-service/internal/repository"
)

type CreateOrder struct {
	orderRepository repository.Order
}

type CreateOrderConfig struct {
	OrderRepository repository.Order
}

type CreateOrderParams struct {
	TicketId, Title, Description, Status string
	TotalPrice                           int
}

func (co *CreateOrder) Execute(params *CreateOrderParams) (*entity.Order, error) {
	return co.orderRepository.CreateOrder(&repository.CreateOrderParams{
		TicketId:    params.TicketId,
		Title:       params.Title,
		Description: params.Description,
		Status:      params.Status,
		TotalPrice:  params.TotalPrice,
	})
}

func NewCreateOrder(config *CreateOrderConfig) CreateOrder {
	return CreateOrder{orderRepository: config.OrderRepository}
}
