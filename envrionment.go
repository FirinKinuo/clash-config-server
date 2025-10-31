package main

import (
	"errors"
	"os"
)

var (
	ErrNoConfigPath = errors.New("no config path")
	ErrNoAddr       = errors.New("no addr provided")
)

type Environment struct {
	ConfigPath string
	Addr       string
}

func LoadEnvironment() (*Environment, error) {
	e := &Environment{}

	e.ConfigPath = os.Getenv("CCS_CONFIG_PATH")
	if e.ConfigPath == "" {
		return nil, ErrNoConfigPath
	}

	e.Addr = os.Getenv("CCS_ADDR")
	if e.Addr == "" {
		return nil, ErrNoAddr
	}

	return e, nil
}
