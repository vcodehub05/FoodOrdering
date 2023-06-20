package app

import (
	"foodApp/internal/config"
	orderService "foodApp/internal/service/order"
	"foodApp/internal/service/order/model"
	orderRepo "foodApp/internal/service/order/repository"
	"foodApp/pkg/messageBroker"

	"go.mongodb.org/mongo-driver/mongo"
)

type services struct {
	orderSvc model.Service
}

type repos struct {
	orderRepo model.Repository
}

func buildServices(cfg *config.Config, db *mongo.Client, messageBroker messageBroker.MessageBroker) *services {
	svc := &services{}
	repo := &repos{}
	repo.buildRepos(db, cfg)

	svc.buildOrderService(cfg, *repo, messageBroker)

	return svc
}

func (r *repos) buildRepos(db *mongo.Client, cfg *config.Config) {
	r.orderRepo = orderRepo.NewRepository(db, *cfg)
}

func (s *services) buildOrderService(cfg *config.Config, repos repos, messageBroker messageBroker.MessageBroker) {
	s.orderSvc = orderService.NewService(messageBroker, repos.orderRepo, *cfg)
}
