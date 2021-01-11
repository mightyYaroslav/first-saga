package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`

	OrderId primitive.ObjectID `bson:"ticket_id" json:"ticket_id,omitempty"`

	Title string `json:"title,omitempty"`

	Status string `json:"status,omitempty"`

	Dishes []string `json:"dishes,omitempty"`
}
