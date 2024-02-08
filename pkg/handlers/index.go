package handlers

import (
	"io"

	"github.com/jontitorr/site/pkg/routes"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct {
}

func (h IndexHandler) Show(c echo.Context) error {
	return render(c, routes.Index())
}

func (h IndexHandler) Save(w io.Writer) error {
	return save(w, routes.Index())
}
