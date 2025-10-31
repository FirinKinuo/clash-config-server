package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

var (
	version = "unknown"
)

func main() {
	slog.Info("starting clash config server", "ver", version)

	envConfig, err := LoadEnvironment()
	if err != nil {
		slog.Error("load environment", "error", err)
		os.Exit(1)
	}

	cfg, err := ReadYamlConfig(envConfig.ConfigPath)
	if err != nil {
		slog.Error("read config", "err", err.Error())
		os.Exit(1)
	}

	router := NewRouter()
	router.BindPaths(cfg.PathRouting)

	slog.Info(fmt.Sprintf("starting listening at: %s", envConfig.Addr))
	err = http.ListenAndServe(envConfig.Addr, router)
	if err != nil {
		slog.Error("start server", "err", err.Error())
		os.Exit(1)
	}
}
