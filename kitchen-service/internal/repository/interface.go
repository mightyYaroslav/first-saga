package repository

import "github.com/mightyYaroslav/first-saga/kitchen-service/internal/entity"

type CreateTicketParams struct {
	Id, OrderId, Title string
	Dishes             []string
}

type Ticket interface {
	Create(req *CreateTicketParams) (*entity.Ticket, error)
	Approve(ticketId string) error
	Reject(ticketId string) error
}
