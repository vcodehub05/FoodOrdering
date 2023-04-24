package customer

import (
	"FoodOrdering/dto"
	"FoodOrdering/model"
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func (s *service) CreateOrder(ctx context.Context, input *dto.CreateOrder) (interface{}, error) {
	order := &model.Order{
		Id:          uuid.NewString(),
		OrderDetail: input.OrderDetail,
		ResturantID: input.ResturantId,
		CustomerID:  input.CustomerId,
		Status:      "inqueue",
	}
	ch, err := s.rabbitMQsvc.CreateChannel()
	if err != nil {
		return nil, err
	}
	queue, err := s.rabbitMQsvc.CreateQueue(ch)
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	}
	err = s.rabbitMQsvc.Publish(ch, queue.Name, msg)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
