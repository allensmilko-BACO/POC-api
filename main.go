package main

import (
	cc "Baco/POC-api/context-components"
	"Baco/POC-api/handlers"
)

func main() {
	views := cc.NewRouterViews()
	dashboard := cc.NewRouterDashboardAPI()
	profile := cc.NewRouterProfileAPI()
	server, serverAddr := handlers.ServerBuilder()
	handlers.Router(server.Group("/api/dashboard"), dashboard)
	handlers.Router(server.Group("/api/profile"), profile)
	handlers.Router(server.Group(""), views)
	server.Logger.Fatal(server.Start(serverAddr))
}
