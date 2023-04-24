package controller

import (
	"FoodOrdering/dto"
	"FoodOrdering/server"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *controller) CreateOrder(c *gin.Context) {
	var ip *dto.CreateOrder
	if err := c.ShouldBindJSON(&ip); err != nil {
		c.JSON(http.StatusUnprocessableEntity, &server.Response{
			Data: err,
		})
		return
	}

	o, err := ctrl.service.CreateOrder(context.TODO(), ip)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, &server.Response{
			Data: err,
		})
		return
	}
	c.JSON(http.StatusOK, &server.Response{
		Data: o,
	})
}
