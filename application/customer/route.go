package main

import (
	controller "FoodOrdering/application/customer/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routes(ctrl controller.API) *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r.POST("/orders", ctrl.CreateOrder)

	return r
}
