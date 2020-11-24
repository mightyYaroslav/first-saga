package repository

import (
	"context"

	"github.com/mightyYaroslav/first-saga/order-service/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderAdapter struct {
	mongoClient *mongo.Client
}

func (o *OrderAdapter) CreateOrder(ticketId, title, description, status string, totalPrice int) (*entity.Order, error) {
	collection := o.mongoClient.Database("example").Collection("order")
	hex, err := primitive.ObjectIDFromHex(ticketId)
	if err != nil {
		return nil, err
	}
	order := entity.Order{
		TicketId:    hex,
		Title:       title,
		Description: description,
		TotalPrice:  totalPrice,
		Status:      status,
	}
	res, err := collection.InsertOne(context.Background(), order)
	if err != nil {
		return nil, err
	}
	order.ID = res.InsertedID.(primitive.ObjectID)
	return &order, nil
}

func (o *OrderAdapter) ApproveOrder(orderId string) error {
	collection := o.mongoClient.Database("example").Collection("order")
	hex, err := primitive.ObjectIDFromHex(orderId)
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

func (o *OrderAdapter) RejectOrder(orderId string) error {
	collection := o.mongoClient.Database("example").Collection("order")
	hex, err := primitive.ObjectIDFromHex(orderId)
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

type OrderRepositoryAdapterConfig struct {
	MongoClient *mongo.Client
}

func NewOrderRepositoryAdapter(config *OrderRepositoryAdapterConfig) Order {
	return &OrderAdapter{mongoClient: config.MongoClient}
}
