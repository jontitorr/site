package handlers

import (
	"github.com/jontitorr/site/pkg/core"
	"github.com/labstack/echo/v4/middleware"
)

func LoadRoutes(s *core.Server) {
	e := s.Echo

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Static("/static", "static")
	e.GET("/", IndexHandler{}.Show)
	e.GET("/projects", ProjectsHandler{}.Show(s))
}
