package messageBroker

import (
	"github.com/streadway/amqp"
)

type MessageBroker interface {
	CreateChannel() (*amqp.Channel, error)
	CreateQueue(ch *amqp.Channel) (*amqp.Queue, error)
	Publish(ch *amqp.Channel, queueName string, msg amqp.Publishing) error
	Consume(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error)
}
