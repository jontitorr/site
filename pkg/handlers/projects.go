package handlers

import (
	"context"
	"io"
	"log/slog"
	"time"

	"github.com/jontitorr/site/pkg/core"
	"github.com/jontitorr/site/pkg/routes"
	"github.com/labstack/echo/v4"
)

type Repositories struct {
	Nodes []struct {
		Name             string
		URL              string
		Description      string
		StargazerCount   int
		ForkCount        int
		PushedAt         time.Time
		RepositoryTopics struct {
			Nodes []struct {
				Topic struct {
					Name string
				}
				URL string
			}
		} `graphql:"repositoryTopics(first: $topics)"`
		Languages struct {
			Nodes []struct {
				Name  string
				Color string
			}
		} `graphql:"languages(first: 1, orderBy: {field: SIZE, direction: DESC})"`
	}
	PageInfo struct {
		EndCursor   string
		HasNextPage bool
	}
}

type ProjectsHandler struct {
}

func (h ProjectsHandler) Show(s *core.Server) (ret func(c echo.Context) error) {
	var err error

	defer func() {
		if err != nil {
			ret = func(c echo.Context) error {
				return err
			}
		}
	}()

	vars, err := s.FetchData(context.Background())
	if err != nil {
		slog.Error("failed to fetch data", slog.Any("error", err))
	}

	ret = func(c echo.Context) error {
		return render(c, routes.Projects(vars.Projects))
	}

	return
}

func (h ProjectsHandler) Save(s *core.Server, w io.Writer) error {
	vars, err := s.FetchData(context.Background())
	if err != nil {
		slog.Error("failed to fetch data", slog.Any("error", err))
		return err
	}

	return save(w, routes.Projects(vars.Projects))
}
