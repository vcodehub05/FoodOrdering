package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"foodApp/internal/config"
	"foodApp/pkg/db"
	"foodApp/pkg/log"
	"foodApp/pkg/messageBroker"
	"foodApp/pkg/messageBroker/rabbitMq"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

const timeout = 5 * time.Second

type Application struct {
	log           log.Logger
	db            *mongo.Client
	messageBroker messageBroker.MessageBroker
	cfg           *config.Config
	router        *mux.Router
	httpServer    *http.Server
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

	client, err := db.ProvideDatabase(log, ctx,
		db.NewConnection(config.DBConfig.User, config.DBConfig.Password), config.DBConfig.Collection, config.DBConfig.Name)
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

	router := mux.NewRouter()
	a.router = router

	services := buildServices(config, a.db, a.messageBroker)
	a.services = services

	a.SetupHandlers()
}

func (a *Application) Start(ctx context.Context) {
	a.httpServer = &http.Server{
		Addr:              ":" + fmt.Sprintf("%v", a.cfg.Server.HTTP.Port),
		Handler:           a.router,
		ReadHeaderTimeout: timeout,
	}
	go func() {
		defer a.log.Infof("server stopped listening")
		if err := a.httpServer.ListenAndServe(); err != nil {
			a.log.Errorf("failed to listen and serve: %v ", err)
			return
		}
	}()
	a.log.Infof("http server started on %d ...", a.cfg.Server.HTTP.Port)
}

func (a *Application) Stop(ctx context.Context) {
	err := a.httpServer.Shutdown(ctx)
	if err != nil {
		a.log.Error(err)
	}

	a.log.Info("shutting down....")
}
