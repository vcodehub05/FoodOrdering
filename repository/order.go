package repository

import (
	"FoodOrdering/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Order interface {
	Create(collection *mongo.Collection, ctx context.Context, order model.Order) (interface{}, error)
	Get(collection *mongo.Collection, ctx context.Context, order model.Order) (*model.Order, error)
}

func initOrder() Order {
	return &order{}
}

type order struct {
	_ struct{}
}

func (o *order) Create(collection *mongo.Collection, ctx context.Context, order model.Order) (interface{}, error) {

	req, err := collection.InsertOne(ctx, order)

	if err != nil {
		fmt.Println("error", err)
		return "", err
	}

	return req.InsertedID, nil
}
func (o *order) Get(collection *mongo.Collection, ctx context.Context, order model.Order) (*model.Order, error) {
	filter := bson.M{"id": order.Id}

	var result model.Order
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil

}
