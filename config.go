package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	PathRouting map[string]Path `yaml:"path-routing"`
}

func ReadYamlConfig(path string) (*Config, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	cfg := &Config{}

	err = yaml.Unmarshal(fileBytes, cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", err)
	}

	return cfg, nil
}

type Path struct {
	File      string    `yaml:"file"`
	BasicAuth BasicAuth `yaml:"basic_auth"`
}

type BasicAuth struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
