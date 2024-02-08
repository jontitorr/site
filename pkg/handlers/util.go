package handlers

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func save(w io.Writer, component templ.Component) error {
	return component.Render(context.Background(), w)
}
