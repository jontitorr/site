package core

import (
	"flag"
	"fmt"
	"os"

	"github.com/peterbourgon/ff/v3"
)

type Config struct {
	ListenAddr string
	GitHub     GitHubConfig
}

type GitHubConfig struct {
	AccessToken string
	User        string
}

func LoadConfig() (Config, error) {
	fs := flag.NewFlagSet("site", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage of %s:\n\n", fs.Name())
		fs.PrintDefaults()
	}

	var (
		listenAddr  *string = fs.String("listen.addr", "0.0.0.0:8080", "listen address")
		githubUser  *string = fs.String("github.user", "jontitorr", "github user")
		githubToken *string = fs.String("github.token", "", "github access token")
	)

	ff.Parse(fs, os.Args[1:], ff.WithEnvVars())

	if *listenAddr == "" {
		return Config{}, fmt.Errorf("listen address is required")
	}

	if *githubUser == "" {
		return Config{}, fmt.Errorf("github user is required")
	}

	if *githubToken == "" {
		return Config{}, fmt.Errorf("github token is required")
	}

	return Config{
		ListenAddr: *listenAddr,
		GitHub: GitHubConfig{
			AccessToken: *githubToken,
			User:        *githubUser,
		},
	}, nil
}
