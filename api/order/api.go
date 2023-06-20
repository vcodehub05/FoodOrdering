package order

import (
	"encoding/json"
	"net/http"

	api "foodApp/api"
	"foodApp/internal/service/order/model"
	"foodApp/pkg/log"

	"github.com/gorilla/mux"
)

type resource struct {
	service model.Service
	log     log.Logger
}

func RegisterHandlers(router *mux.Router, service model.Service, log log.Logger) {
	res := resource{
		service: service,
		log:     log,
	}

	orderRouter := router.PathPrefix("/order").Subrouter()
	api.RegisterHandler(orderRouter, "POST", "", nil, res.CreateOrder)
	api.RegisterHandler(orderRouter, "OPTIONS", "", nil, res.CreateOrder)
}

func (res resource) CreateOrder(w http.ResponseWriter, r *http.Request) {
	req := &model.CreateOrder{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(req); err != nil {
		return
	}
	order := model.CreateOrder{
		OrderDetail:  req.OrderDetail,
		CustomerID:   req.CustomerID,
		RestaurantID: req.RestaurantID,
	}

	err := res.service.Create(r.Context(), res.log, order)
	if err != nil {
		api.Write(w, http.StatusInternalServerError, api.NewResponse(false, "failed to place order", err))
		return
	}
	api.Write(w, http.StatusOK, api.NewResponse(true, "Order has been placed", nil))
}
