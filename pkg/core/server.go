package core

import (
	"github.com/labstack/echo/v4"
	"github.com/shurcooL/githubv4"
)

type Server struct {
	Config Config
	Echo   *echo.Echo
	GitHub *githubv4.Client
}

func NewServer(cfg Config, gh *githubv4.Client) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	return &Server{
		Config: cfg,
		Echo:   e,
		GitHub: gh,
	}
}

func (s *Server) Start() error {
	return s.Echo.Start(s.Config.ListenAddr)
}
