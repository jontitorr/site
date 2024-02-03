package core

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ListenAddr string       `yaml:"listen_addr"`
	GitHub     GitHubConfig `yaml:"github"`
}

type GitHubConfig struct {
	AccessToken string `yaml:"access_token"`
	User        string `yaml:"user"`
}

func LoadConfig(path string) (cfg Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}

	if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return Config{}, err
	}
	return
}
