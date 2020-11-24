package usecase

import (
	"github.com/mightyYaroslav/first-saga/order-service/internal/repository"
)

type ApproveTicket struct {
	ticketRepository repository.Ticket
}

type ApproveTicketConfig struct {
	TicketRepository repository.Ticket
}

func (ao *ApproveTicket) Execute(ticketId string) error {
	return ao.ticketRepository.Approve(ticketId)
}

func NewApproveOrder(config *ApproveTicketConfig) ApproveTicket {
	return ApproveTicket{ticketRepository: config.TicketRepository}
}
