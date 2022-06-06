package contextComponents

import (
	apiControllers "Baco/POC-api/api/controllers"
	"Baco/POC-api/entities"
)

func NewRouterDashboardAPI() []entities.RouteHandler {
	var routerDashboard []entities.RouteHandler

	var dashboard entities.RouteHandler
	dashboard.Path = "/test/:t"
	dashboard.Method = "GET"
	dashboard.Function = apiControllers.GetDashboardAPI

	routerDashboard = append(routerDashboard, dashboard)

	return routerDashboard
}
