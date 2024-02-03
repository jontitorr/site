package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/jontitorr/site/pkg/core"
	"github.com/jontitorr/site/pkg/handlers"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func main() {
	cfgPath := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	cfg, err := core.LoadConfig(*cfgPath)
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		os.Exit(-1)
	}

	githubClient := githubv4.NewClient(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GitHub.AccessToken},
	)))

	s := core.NewServer(cfg, githubClient)

	handlers.LoadRoutes(s)

	go s.Start()

	slog.Info("started jontitor.com", slog.Any("url", fmt.Sprintf("http://%s", cfg.ListenAddr)))
	si := make(chan os.Signal, 1)
	signal.Notify(si, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-si
}
