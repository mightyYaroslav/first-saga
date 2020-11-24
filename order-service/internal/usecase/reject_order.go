package usecase

import (
	"github.com/mightyYaroslav/first-saga/order-service/internal/repository"
)

type RejectOrder struct {
	orderRepository repository.Order
}

type RejectOrderConfig struct {
	OrderRepository repository.Order
}

func (ro *RejectOrder) Execute(orderId string) error {
	return ro.orderRepository.RejectOrder(orderId)
}

func NewRejectOrder(config *RejectOrderConfig) RejectOrder {
	return RejectOrder{orderRepository: config.OrderRepository}
}
