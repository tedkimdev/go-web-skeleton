package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/tedkimdev/go-web-skeleton/internal/server/httpapi/router"
	"github.com/tedkimdev/go-web-skeleton/internal/service"
)

type Server struct {
	HttpServer *http.Server
}

func New(config *Config, db *sqlx.DB, service *service.Service) *Server {
	if config == nil {
		panic("config is nil")
	}
	if db == nil {
		panic("db is nil")
	}
	if service == nil {
		panic("service is nil")
	}

	httpRouter := router.NewHttpRouter(db, service)

	httpServer := http.Server{
		Addr:    config.Addr,
		Handler: httpRouter,
	}

	return &Server{
		HttpServer: &httpServer,
	}
}

func (s *Server) Run() error {
	slog.Info(fmt.Sprintf("server started on %s", s.HttpServer.Addr))
	return s.HttpServer.ListenAndServe()
}
