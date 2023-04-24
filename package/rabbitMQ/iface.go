package rabbitmq

import (
	"github.com/streadway/amqp"
)

type API interface {
	CreateChannel() (*amqp.Channel, error)
	CreateQueue(ch *amqp.Channel) (*amqp.Queue, error)
	Publish(ch *amqp.Channel, queueName string, msg amqp.Publishing) error
	Consume(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error)
}

type rabbitmq struct {
	conn *amqp.Connection
}

func New(conn *amqp.Connection) API {
	return &rabbitmq{
		conn:conn,
	}

}
