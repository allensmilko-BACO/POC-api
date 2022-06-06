package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProfileAPI(c echo.Context) error {
	id := c.Param("t")
	return c.String(http.StatusOK, id)
}
