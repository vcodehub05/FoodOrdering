package main

import (
	"FoodOrdering/internal/mongo"
	"FoodOrdering/model"
	rabbitmq "FoodOrdering/package/rabbitMQ"
	"FoodOrdering/repository"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false})
}
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
		log.Err(fmt.Errorf("failed to create channel %w", err))
		return
	}
	q, err := rabbitMQsvc.CreateQueue(ch)
	if err != nil {
		log.Err(fmt.Errorf("failed to create queue %w", err))
		return
	}
	prefetchCount := 2
	err = ch.Qos(prefetchCount, 0, false)
	if err != nil {
		log.Err(fmt.Errorf("Failed to set QoS: %w", err))
	}
	numMsgs := q.Messages
	// Launch multiple consumers to handle messages
	numConsumers := numMsgs / prefetchCount
	if numMsgs <= 1 {
		numConsumers = 1
	}
	for i := 0; i < numConsumers; i++ {
		go func(workerNum int) {
			// Consume messages from queue
			messages, err := rabbitMQsvc.Consume(ch, q.Name)
			if err != nil {
				log.Err(fmt.Errorf("Failed to consume messages from the queue: %v", err))
			}
			for message := range messages {
				var order model.Order
				json.Unmarshal(message.Body, &order)
				order.Status = "rececived"
				_, err := mongosvc.Order().Create(collection, ctx, order)
				if err != nil {
					log.Err(fmt.Errorf("Failed to insert in db: %v", err))
				}
			}
		}(i)
	}
	// Wait for goroutines to finish
	select {}
}
