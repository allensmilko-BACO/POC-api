package entities

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type RouteHandler struct {
	Path     string
	Method   string
	Function echo.HandlerFunc
}

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

type PageInfo struct {
	Title       string
	Description string
}
