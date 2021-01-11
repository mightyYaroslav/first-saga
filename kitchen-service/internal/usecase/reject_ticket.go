package usecase

import (
	"github.com/mightyYaroslav/first-saga/kitchen-service/internal/repository"
)

type RejectTicket struct {
	ticketRepository repository.Ticket
}

type RejectTicketConfig struct {
	TicketRepository repository.Ticket
}

func (ro *RejectTicket) Execute(ticketId string) error {
	return ro.ticketRepository.Reject(ticketId)
}

func NewRejectTicket(config *RejectTicketConfig) RejectTicket {
	return RejectTicket{ticketRepository: config.TicketRepository}
}
