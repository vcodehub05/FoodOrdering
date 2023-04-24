package main

import (
	"FoodOrdering/internal/mongo"
	"FoodOrdering/model"
	rabbitmq "FoodOrdering/package/rabbitMQ"
	"FoodOrdering/repository"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()
	opt := mongo.CreateMongoOption("vivekk", "UqqVDqz7OPMgYiwn")
	client, err := mongo.ProvideDatabase(ctx, opt)
	collection := client.Database("foodApp").Collection("order")
	if err != nil {
		return
	}
	mongosvc := repository.New(client)
	conn, err := rabbitmq.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return
	}
	rabbitMQsvc := rabbitmq.New(conn)

	ch, err := rabbitMQsvc.CreateChannel()
	if err != nil {
		fmt.Println("error", err)
		return
	}
	q, err := rabbitMQsvc.CreateQueue(ch)
	if err != nil {
		return
	}
	// Consume messages from the queue
	messages, err := rabbitMQsvc.Consume(ch, q.Name)
	if err != nil {
		log.Fatalf("Failed to consume messages from the queue: %s", err)
	}
	for message := range messages {
		var order model.Order
		json.Unmarshal(message.Body, &order)
		order.Status = "rececived"
		_, err := mongosvc.Order().Create(collection, ctx, order)
		if err != nil {
			fmt.Printf("err%v", err)
		}
	}

}
