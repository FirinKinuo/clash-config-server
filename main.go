package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gopkg.in/yaml.v3"
)

var (
	version = "unknown"
)

type Flags struct {
	ConfigPath string
}

func newFlags() Flags {
	flags := Flags{}

	flag.StringVar(&flags.ConfigPath, "c", "config.yaml", "config file path")

	flag.Parse()

	return flags
}

func readConfig(path string) (Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read file: %w", err)
	}

	var c Config

	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return Config{}, fmt.Errorf("unmarshal yaml: %w", err)
	}

	return c, nil
}

func serveFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	}
}

func createRouter(cfg Config) chi.Router {
	rootRouter := chi.NewRouter()

	rootRouter.Use(middleware.RequestID)
	rootRouter.Use(middleware.RealIP)
	rootRouter.Use(middleware.Logger)
	rootRouter.Use(middleware.Recoverer)

	for key, extraPath := range cfg.ExtraPaths {
		basicAuth := map[string]string{extraPath.BasicAuth.User: extraPath.BasicAuth.Password}

		extraRouter := chi.NewRouter()
		extraRouter.Use(middleware.BasicAuth(key, basicAuth))
		extraRouter.Get("/", serveFile(extraPath.Path))

		rootRouter.Mount("/"+key, extraRouter)
	}

	return rootRouter
}

func main() {
	slog.Info("starting server", "ver", version)
	flags := newFlags()
	cfg, err := readConfig(flags.ConfigPath)
	if err != nil {
		slog.Error("read config", "err", err.Error())
		return
	}

	router := createRouter(cfg)

	slog.Info("listen and serve at 0.0.0.0:8000")
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		slog.Error("start server", "err", err.Error())
		return
	}
}
