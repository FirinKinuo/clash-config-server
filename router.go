package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	*chi.Mux
}

func NewRouter() *Router {
	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)

	return &Router{Mux: r}
}

func (r *Router) BindPaths(pathRouting map[string]Path) {
	for route, path := range pathRouting {
		pathRouter := chi.NewRouter()

		if path.BasicAuth != (BasicAuth{}) {
			basicAuth := map[string]string{path.BasicAuth.User: path.BasicAuth.Password}
			pathRouter.Use(middleware.BasicAuth(route, basicAuth))
		}

		pathRouter.Get("/", ServeFile(path.File))

		r.Mount("/"+route, pathRouter)
	}
}
