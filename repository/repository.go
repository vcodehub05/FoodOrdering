package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type API interface {
	Order() Order
}

func New(conn *mongo.Client) API {
	return &api{
		conn:  conn,
		order: initOrder(),
	}

}

type api struct {
	conn  *mongo.Client
	order Order
}

func (a *api) Order() Order {
	return a.order
}
