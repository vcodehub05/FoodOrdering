package app

import (
	"context"
	"time"

	"foodApp/internal/config"
	"foodApp/pkg/db"
	"foodApp/pkg/log"
	"foodApp/pkg/messageBroker"
	"foodApp/pkg/messageBroker/rabbitMq"

	"go.mongodb.org/mongo-driver/mongo"
)

const timeout = 5 * time.Second

type Application struct {
	log           log.Logger
	db            *mongo.Client
	messageBroker messageBroker.MessageBroker
	cfg           *config.Config
	services      *services
}

func (a *Application) Init(ctx context.Context, configFile string) {
	log := log.New().With(ctx)
	a.log = log

	config, err := config.Load(configFile)
	if err != nil {
		log.Fatalf("failed to read config: %s ", err)
		return
	}
	a.cfg = config

	client, err := db.ProvideDatabase(log, ctx, db.NewConnection(config.DBConfig.User, config.DBConfig.Password), config.DBConfig.Collection, config.DBConfig.Name)
	if err != nil {
		log.Fatalf("error connecting db: %s ", err)
		return
	}
	a.db = client

	rabbitmqConn, err := rabbitMq.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return
	}
	rabbitmqsvc := rabbitMq.New(rabbitmqConn)
	a.messageBroker = rabbitmqsvc

	services := buildServices(config, a.db, a.messageBroker)
	a.services = services
}

func (a *Application) Start(ctx context.Context) {
	err := a.services.orderSvc.Receive(ctx, a.log)
	if err != nil {
		return
	}
}
