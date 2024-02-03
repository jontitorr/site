package handlers

import (
	"github.com/jontitorr/site/pkg/routes"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct {
}

func (h IndexHandler) Show(c echo.Context) error {
	return render(c, routes.Index())
}
