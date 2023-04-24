package dto

import "FoodOrdering/model"

type CreateOrder struct {
	CustomerId  string            `json:"customer_id"`
	ResturantId string            `json:"resturant_id"`
	OrderDetail model.OrderDetail `json:"order_detail"`
}
