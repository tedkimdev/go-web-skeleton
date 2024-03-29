package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/tedkimdev/go-web-skeleton/internal/repository"
	"github.com/tedkimdev/go-web-skeleton/internal/server"
	"github.com/tedkimdev/go-web-skeleton/internal/server/database"
	"github.com/tedkimdev/go-web-skeleton/internal/service"
)

func main() {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(jsonHandler)
	slog.SetDefault(logger)

	config, err := server.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(config.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := repository.NewRepoistory(db)
	service := service.NewService(repository)

	server := server.New(config, db, service)
	if server == nil {
		log.Fatal("server is nil")
	}

	err = server.Run()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
