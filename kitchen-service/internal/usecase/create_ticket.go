package usecase

import (
	"github.com/mightyYaroslav/first-saga/kitchen-service/internal/repository"
)

type CreateTicket struct {
	ticketRepository repository.Ticket
}

type CreateTicketConfig struct {
	TicketRepository repository.Ticket
}

type CreateTicketParams struct {
	Id, OrderId, Title string
	Dishes             []string
}

func (ao *CreateTicket) Execute(params *CreateTicketParams) error {
	_, err := ao.ticketRepository.Create(
		&repository.CreateTicketParams{
			Id:      params.Id,
			OrderId: params.OrderId,
			Title:   params.Title,
			Dishes:  params.Dishes,
		},
	)
	return err
}

func NewCreateTicket(config *CreateTicketConfig) CreateTicket {
	return CreateTicket{ticketRepository: config.TicketRepository}
}
