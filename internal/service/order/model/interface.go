package model

import (
	"context"

	"foodApp/pkg/log"
)

type Service interface {
	Create(ctx context.Context, log log.Logger, input CreateOrder) error
	Receive(ctx context.Context, log log.Logger) error
}

type Repository interface {
	Add(ctx context.Context, log log.Logger, input Order) error
}
