package controller

import (
	"FoodOrdering/service/customer"

	"github.com/gin-gonic/gin"
)

type API interface {
	CreateOrder(c *gin.Context)
}

func New(s customer.Service) API {
	return &controller{
		service: s,
	}
}

type controller struct {
	service customer.Service
}
