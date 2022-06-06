package api

import (
	"net/http"

	"Baco/POC-api/entities"

	"github.com/labstack/echo/v4"
)

func GetHome(c echo.Context) error {
	contextPage := entities.PageInfo{Title: "Home", Description: "Loremn ipsum"}

	return c.Render(http.StatusOK, "home.html", contextPage)

}

func GetAuth(c echo.Context) error {
	contextPage := entities.PageInfo{Title: "Registro", Description: "Loremn ipsum"}

	return c.Render(http.StatusOK, "auth.html", contextPage)

}
