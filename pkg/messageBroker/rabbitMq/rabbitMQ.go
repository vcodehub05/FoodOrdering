package rabbitMq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Rabbitmq struct {
	conn *amqp.Connection
}

func New(conn *amqp.Connection) *Rabbitmq {
	return &Rabbitmq{
		conn: conn,
	}
}

func Dial(url string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	return conn, nil
}

func (r *Rabbitmq) CreateChannel() (*amqp.Channel, error) {
	channelRabbitMQ, err := r.conn.Channel()
	if err != nil {
		return nil, err
	}
	return channelRabbitMQ, nil
}

func (r *Rabbitmq) CreateQueue(ch *amqp.Channel) (*amqp.Queue, error) {
	queue, err := ch.QueueDeclare(
		"QueueService1", // queue name
		true,            // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		return nil, err
	}
	return &queue, nil
}

func (r *Rabbitmq) Publish(ch *amqp.Channel, queueName string, msg amqp.Publishing) error {
	if err := ch.Publish(
		"",        // exchange
		queueName, // queue name
		false,     // mandatory
		false,     // immediate
		msg,       // message to publish
	); err != nil {
		return err
	}
	return nil
}

func (r *Rabbitmq) Consume(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	messages, err := ch.Consume(
		"QueueService1", // queue name
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no local
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
