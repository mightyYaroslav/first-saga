package usecase

import (
	"github.com/mightyYaroslav/first-saga/order-service/internal/repository"
)

type ApproveOrder struct {
	orderRepository repository.Order
}

type ApproveOrderConfig struct {
	OrderRepository repository.Order
}

func (ao *ApproveOrder) Execute(orderId string) error {
	return ao.orderRepository.ApproveOrder(orderId)
}

func NewApproveOrder(config *ApproveOrderConfig) ApproveOrder {
	return ApproveOrder{orderRepository: config.OrderRepository}
}
