package repository

type Ticket interface {
	Approve(ticketId string) error
	Reject(ticketId string) error
}
