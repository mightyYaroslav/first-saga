package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`

	TicketId primitive.ObjectID `bson:"ticket_id" json:"ticket_id,omitempty"`

	Title string `json:"title,omitempty"`

	Description string `json:"description,omitempty"`

	TotalPrice int `json:"total_price,omitempty"`

	Status string `json:"status,omitempty"`
}
