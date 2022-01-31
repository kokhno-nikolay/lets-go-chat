package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	api "github.com/kokhno-nikolay/lets-go-chat/internal/api/http"
	"github.com/kokhno-nikolay/lets-go-chat/internal/config"
	"github.com/kokhno-nikolay/lets-go-chat/internal/repository"
	"github.com/kokhno-nikolay/lets-go-chat/internal/repository/postgres"
	"github.com/kokhno-nikolay/lets-go-chat/internal/server"
	"github.com/kokhno-nikolay/lets-go-chat/internal/service"
)

func main() {
	cfg := config.GetConfig()
	ctx := context.Background()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	dbClient, err := postgres.NewClient("postgres", cfg)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	repos := repository.NewRepositories(dbClient)
	services := service.NewServices(service.Deps{Repos: repos})
	handlers := api.NewHandler(services)

	srv := server.NewServer(cfg, handlers.Init(cfg))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	log.Printf("Server started")

	<-quit

	if err := srv.Stop(ctx); err != nil {
		log.Printf("failed to stop server: %v", err.Error())
	}
}
