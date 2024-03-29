package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/tedkimdev/go-web-skeleton/internal/server/httpapi/ping"
	"github.com/tedkimdev/go-web-skeleton/internal/server/httpapi/posts"
	"github.com/tedkimdev/go-web-skeleton/internal/service"
)

func NewHttpRouter(db *sqlx.DB, service *service.Service) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", ping.GetPing)
	router.Route("/api/v1", func(router chi.Router) {
		router.Get("/ping", ping.GetPing)
		posts.Route(router, service.PostService)
	})

	return router
}
