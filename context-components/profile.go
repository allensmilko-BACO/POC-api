package contextComponents

import (
	apiControllers "Baco/POC-api/api/controllers"
	"Baco/POC-api/entities"
)

func NewRouterProfileAPI() []entities.RouteHandler {
	var routerProfile []entities.RouteHandler

	var profile entities.RouteHandler
	profile.Path = "test"
	profile.Method = "GET"
	profile.Function = apiControllers.GetProfileAPI

	routerProfile = append(routerProfile, profile)

	return routerProfile
}
