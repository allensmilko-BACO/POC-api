package contextComponents

import (
	apiControllers "Baco/POC-api/api/controllers"
	"Baco/POC-api/entities"
)

func NewRouterViews() []entities.RouteHandler {
	var routerViews []entities.RouteHandler

	var home entities.RouteHandler
	home.Path = "/"
	home.Method = "GET"
	home.Function = apiControllers.GetHome

	routerViews = append(routerViews, home)

	var auth entities.RouteHandler
	auth.Path = "/auth"
	auth.Method = "GET"
	auth.Function = apiControllers.GetAuth

	routerViews = append(routerViews, auth)

	return routerViews
}
