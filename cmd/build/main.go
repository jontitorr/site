package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"path"

	"github.com/jontitorr/site/pkg/core"
	"github.com/jontitorr/site/pkg/handlers"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func main() {
	cfg, err := core.LoadConfig()
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		os.Exit(-1)
	}

	githubClient := githubv4.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GitHub.AccessToken},
	)))

	s := core.NewServer(cfg, githubClient)

	rootPath := "public"
	if err := os.Mkdir(rootPath, 0755); err != nil && !os.IsExist(err) {
		log.Fatalf("failed to create output directory: %v", err)
	}

	for _, route := range handlers.Routes(s) {
		name := path.Join(rootPath, route.Name+".html")
		f, err := os.Create(name)
		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}

		err = route.Save(f)
		if err != nil {
			log.Fatalf("failed to render route: %v", err)
		}
	}
}
