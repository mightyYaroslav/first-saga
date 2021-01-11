package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Consumer struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`

	FirstName string `json:"first_name,omitempty"`

	LastName string `json:"last_name,omitempty"`

	Phone string `json:"phone,omitempty"`
}
