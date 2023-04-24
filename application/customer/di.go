package main

import (
	"FoodOrdering/application/customer/controller"
	"FoodOrdering/internal/mongo"
	rabbitmq "FoodOrdering/package/rabbitMQ"
	"FoodOrdering/repository"
	"FoodOrdering/server"
	"FoodOrdering/service/customer"
	"context"
	"net/http"
)

func di(port int) (*http.Server, error) {
	ctx := context.Background()
	opt := mongo.CreateMongoOption("vivekk", "UqqVDqz7OPMgYiwn")
	client, err := mongo.ProvideDatabase(ctx, opt)
	if err != nil {
		return nil, err
	}
	collection := client.Database("foodApp").Collection("order")
	repository := repository.New(client)
	rabbitmqConn, err := rabbitmq.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	rabbitmqsvc := rabbitmq.New(rabbitmqConn)
	service := customer.New(&customer.InitialiseServiceInput{
		Repo:        repository,
		Collection:  collection,
		RabbitMQsvc: rabbitmqsvc,
	})
	ctrl := controller.New(service)
	r := routes(ctrl)
	s := server.New(r, port)
	return s, nil
}
