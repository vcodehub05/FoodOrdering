package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ProvideDatabase(ctx context.Context, option *options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, option)
	if err != nil {
		return nil, err
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	return client, nil
}

func CreateMongoOption(DbUser string, DbPassword string) *options.ClientOptions {
	return options.Client().ApplyURI("mongodb+srv://" + DbUser + ":" + DbPassword + "@cluster0.oldl3xt.mongodb.net/?retryWrites=true&w=majority")
}
