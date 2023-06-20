package app

import (
	api "foodApp/api/order"
)

func (a *Application) SetupHandlers() {
	api.RegisterHandlers(a.router, a.services.orderSvc, a.log)
}
