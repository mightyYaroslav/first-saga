package repository

import (
	"context"

	"github.com/mightyYaroslav/first-saga/kitchen-service/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TicketAdapter struct {
	mongoClient *mongo.Client
}

func (o *TicketAdapter) Create(req *CreateTicketParams) (*entity.Ticket, error) {
	collection := o.mongoClient.Database("example").Collection("ticket")
	hexId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}
	hexOrderId, err := primitive.ObjectIDFromHex(req.OrderId)
	if err != nil {
		return nil, err
	}
	order := entity.Ticket{
		ID:      hexId,
		OrderId: hexOrderId,
		Title:   req.Title,
		Dishes:  req.Dishes,
		Status:  "pending",
	}
	res, err := collection.InsertOne(context.Background(), order)
	if err != nil {
		return nil, err
	}
	order.ID = res.InsertedID.(primitive.ObjectID)
	return &order, nil
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
