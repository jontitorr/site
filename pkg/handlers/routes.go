package handlers

import (
	"io"

	"github.com/jontitorr/site/pkg/core"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoadRoutes(s *core.Server) {
	e := s.Echo

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Static("/static", "public/static")

	for _, r := range Routes(s) {
		loadRoute(s, r)
	}
}

type Route struct {
	Name   string
	Method string
	Url    string
	Show   echo.HandlerFunc
	Save   func(io.Writer) error
}

func loadRoute(s *core.Server, r Route) {
	s.Echo.Add(r.Method, r.Url, r.Show)
}

func Routes(s *core.Server) []Route {
	return []Route{
		{
			Name:   "index",
			Method: "GET",
			Url:    "/",
			Show:   IndexHandler{}.Show,
			Save:   IndexHandler{}.Save,
		},
		{
			Name:   "projects",
			Method: "GET",
			Url:    "/projects",
			Show:   ProjectsHandler{}.Show(s),
			Save:   func(w io.Writer) error { return ProjectsHandler{}.Save(s, w) },
		},
	}
}
