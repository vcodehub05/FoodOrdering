package order

import (
	"context"
	"encoding/json"

	"foodApp/internal/config"
	"foodApp/internal/service/order/model"
	"foodApp/pkg/log"
	"foodApp/pkg/messageBroker"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type Service struct {
	messageBroker messageBroker.MessageBroker
	repo          model.Repository
	config        config.Config
}

func NewService(
	messageBroker messageBroker.MessageBroker,
	repo model.Repository,
	config config.Config,
) *Service {
	return &Service{messageBroker: messageBroker, repo: repo, config: config}
}

func (s *Service) Create(ctx context.Context, log log.Logger, input model.CreateOrder) error {
	order := &model.Order{
		Id:           uuid.NewString(),
		OrderDetail:  input.OrderDetail,
		RestaurantID: input.RestaurantID,
		CustomerID:   input.CustomerID,
		Status:       "inqueue",
	}
	ch, err := s.messageBroker.CreateChannel()
	if err != nil {
		return err
	}
	queue, err := s.messageBroker.CreateQueue(ch)
	if err != nil {
		return err
	}
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	}
	err = s.messageBroker.Publish(ch, queue.Name, msg)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Receive(ctx context.Context, log log.Logger) error {
	ch, err := s.messageBroker.CreateChannel()
	if err != nil {
		log.Errorf("create channel: %w", err)
		return err
	}
	q, err := s.messageBroker.CreateQueue(ch)
	if err != nil {
		return err
	}
	prefetchCount := s.config.RabbitMqConfig.PrefetchCount
	err = ch.Qos(prefetchCount, 0, false)
	if err != nil {
		return err
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
			messages, err := s.messageBroker.Consume(ch, q.Name)
			if err != nil {
				log.Errorf("consuming message: %w", err)
			}
			for message := range messages {
				var order model.Order
				err := json.Unmarshal(message.Body, &order)
				if err != nil {
					log.Errorf("unmarshal body: %w", err)
				}
				order.Status = "rececived"
				err = s.repo.Add(ctx, log, order)
				if err != nil {
					log.Errorf("Inserting in Db %w", err)
				}
			}
		}(i)
	}
	// Wait for goroutines to finish
	select {}
}
