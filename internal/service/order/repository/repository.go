package repo

import (
	"context"

	"foodApp/internal/config"
	"foodApp/internal/service/order/model"
	"foodApp/pkg/log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db     *mongo.Client
	config config.Config
}

func NewRepository(db *mongo.Client, config config.Config) *Repository {
	return &Repository{db: db, config: config}
}

func (r *Repository) Add(ctx context.Context, log log.Logger, order model.Order) error {
	collection := r.db.Database(r.config.DBConfig.Name).Collection(r.config.DBConfig.Collection)
	_, err := collection.InsertOne(ctx, order)
	if err != nil {
		log.Errorf("Insert into db %v", err)
		return err
	}

	return nil
}
