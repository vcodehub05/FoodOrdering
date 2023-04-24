package customer

import (
	"FoodOrdering/dto"
	rabbitmq "FoodOrdering/package/rabbitMQ"
	"FoodOrdering/repository"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	CreateOrder(ctx context.Context, input *dto.CreateOrder) (interface{}, error)
}

type InitialiseServiceInput struct {
	Repo        repository.API
	Collection  *mongo.Collection
	RabbitMQsvc rabbitmq.API
}

func New(input *InitialiseServiceInput) Service {
	return &service{
		repo:        input.Repo,
		collection:  input.Collection,
		rabbitMQsvc: input.RabbitMQsvc,
	}
}

type service struct {
	repo        repository.API
	collection  *mongo.Collection
	rabbitMQsvc rabbitmq.API
}
