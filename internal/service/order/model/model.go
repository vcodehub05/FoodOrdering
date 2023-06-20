package model

import (
	"time"
)

type CreateOrder struct {
	CustomerID   string      `json:"customer_id"`
	RestaurantID string      `json:"restaurant_id"`
	OrderDetail  OrderDetail `json:"order_detail"`
}
type Order struct {
	Id           string      `json:"item_name" bson:"item_name"`
	OrderDetail  OrderDetail `json:"order_detail" bson:"order_detail"`
	CustomerID   string      `json:"customer_id" bson:"customer_id"`
	RestaurantID string      `json:"restaurant_id" bson:"restaurant_id"`
	Status       string      `json:"status" bson:"status"`
	CreatedAt    time.Time   `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at" bson:"updated_at"`
}

type OrderDetail struct {
	ItemName string `json:"item_name" bson:"item_name"`
	Quantity string `json:"quantity" bson:"quantity"`
	Price    uint64 `json:"price" bson:"price"`
}
