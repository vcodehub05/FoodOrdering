package db

import (
	"context"

	"foodApp/pkg/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Name       string `yaml:"name"`
	Collection string `yaml:"collection"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Uri        string `yaml:"uri"`
}

func NewConnection(DbUser string, DbPassword string) *options.ClientOptions {
	return options.Client().ApplyURI("mongodb+srv://" + DbUser + ":" +
		DbPassword + "@cluster0.oldl3xt.mongodb.net/?retryWrites=true&w=majority")
}

func ProvideDatabase(log log.Logger, ctx context.Context,
	option *options.ClientOptions, collection string, database string,
) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, option)
	if err != nil {
		return nil, err
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	db := client.Database(database)
	collectionExists, err := checkCollectionExists(ctx, db, collection)
	if err != nil {
		log.Errorf("check collection %v", err)
	}

	// Create the collection if it doesn't exist
	if !collectionExists {
		err = createCollection(ctx, db, collection)
		if err != nil {
			log.Errorf("create collection %v", err)
		}
	}

	return client, nil
}

func checkCollectionExists(ctx context.Context, db *mongo.Database, collectionName string) (bool, error) {
	collections, err := db.ListCollectionNames(ctx, bson.M{"name": collectionName})
	if err != nil {
		return false, err
	}

	return len(collections) > 0, nil
}

// Function to create a collection
func createCollection(ctx context.Context, db *mongo.Database, collectionName string) error {
	err := db.CreateCollection(ctx, collectionName)
	if err != nil {
		return err
	}

	return nil
}
