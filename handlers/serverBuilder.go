package handlers

import (
	"fmt"
	"html/template"
	"os"

	"Baco/POC-api/entities"

	"github.com/labstack/echo/v4"
)

func ServerBuilder() (*echo.Echo, string) {
	host := "0.0.0.0"
	port := os.Getenv("PORT")
	server := echo.New()
	serverAddr := fmt.Sprintf("%s:%s", host, port)
	t := &entities.Template{
		Templates: template.Must(template.ParseGlob("views/**/*.html")),
	}

	server.Renderer = t
	server.Static("/assets", "views/assets")
	return server, serverAddr
}
