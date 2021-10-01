package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//MiddlewarePipeline define the middleware pipeline
func MiddlewarePipeline(router *chi.Mux) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
}
