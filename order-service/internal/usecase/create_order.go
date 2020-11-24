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

func (co *CreateOrder) Execute(ticketId, title, description, status string, totalPrice int) (*entity.Order, error) {
	return co.orderRepository.CreateOrder(ticketId, title, description, status, totalPrice)
}

func NewCreateOrder(config *CreateOrderConfig) CreateOrder {
	return CreateOrder{orderRepository: config.OrderRepository}
}
