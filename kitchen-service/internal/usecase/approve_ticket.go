package usecase

import (
	"github.com/mightyYaroslav/first-saga/kitchen-service/internal/repository"
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

func NewApproveTicket(config *ApproveTicketConfig) ApproveTicket {
	return ApproveTicket{ticketRepository: config.TicketRepository}
}
