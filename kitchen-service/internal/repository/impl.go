package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TicketAdapter struct {
	mongoClient *mongo.Client
}

func (o *TicketAdapter) Approve(ticketId string) error {
	collection := o.mongoClient.Database("example").Collection("ticket")
	hex, err := primitive.ObjectIDFromHex(ticketId)
	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(
		context.Background(),
		struct{ _id primitive.ObjectID }{_id: hex},
		struct{ status string }{status: "Approved"},
	)
	return err
}

func (o *TicketAdapter) Reject(ticketId string) error {
	collection := o.mongoClient.Database("example").Collection("ticket")
	hex, err := primitive.ObjectIDFromHex(ticketId)
	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(
		context.Background(),
		struct{ _id primitive.ObjectID }{_id: hex},
		struct{ status string }{status: "Rejected"},
	)
	return err
}

type TicketRepositoryAdapterConfig struct {
	MongoClient *mongo.Client
}

func NewTicketRepositoryAdapter(config *TicketRepositoryAdapterConfig) Ticket {
	return &TicketAdapter{mongoClient: config.MongoClient}
}
