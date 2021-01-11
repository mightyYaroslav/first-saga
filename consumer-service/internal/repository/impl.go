package repository

import (
	"context"

	"github.com/mightyYaroslav/first-saga/consumer-service/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConsumerAdapter struct {
	mongoClient *mongo.Client
}

func (c *ConsumerAdapter) Verify(params *ConsumerVerifyParams) error {
	collection := c.mongoClient.Database("example").Collection("consumer")
	result := collection.FindOne(context.Background(), struct {
		first_name string
		last_name  string
		phone      string
	}{
		first_name: params.FirstName,
		last_name:  params.LastName,
		phone:      params.Phone,
	})
	err := result.Err()
	if err != nil {
		return err
	}
	consumer := &entity.Consumer{}
	err = result.Decode(consumer)
	return err
}

type ConsumerRepositoryAdapterConfig struct {
	MongoClient *mongo.Client
}

func NewConsumerRepositoryAdapter(config *ConsumerRepositoryAdapterConfig) Consumer {
	return &ConsumerAdapter{mongoClient: config.MongoClient}
}
