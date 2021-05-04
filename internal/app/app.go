package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/config"
	"github.com/JesusG2000/hexsatisfaction/internal/handler"
	"github.com/JesusG2000/hexsatisfaction/internal/repository"
	"github.com/JesusG2000/hexsatisfaction/internal/server"
	"github.com/JesusG2000/hexsatisfaction/internal/service"
	"github.com/JesusG2000/hexsatisfaction/pkg/auth"
	"github.com/JesusG2000/hexsatisfaction/pkg/database/pg"
)

func Run(configPath string) {

	ctx := context.Background()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal("Init config error: ", err)
	}

	db, err := pg.NewPg(cfg.Pg)
	if err != nil {
		log.Fatal("Init db error: ", err)
	}

	tokenManager, err := auth.NewManager(cfg.Auth.SigningKey)
	if err != nil {
		log.Fatal("Init jwt-token error: ", err)
	}

	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos:        repos,
		TokenManager: tokenManager,
	})

	newHandler := handler.NewHandler(services, tokenManager)
	srv := server.NewServer(cfg, newHandler)

	go startService(ctx, srv)

	<-stop

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Fatal("failed to stop server: ", err)
	}

	log.Printf("shutting down server...")
}

func startService(ctx context.Context, coreService *server.Server) {
	if err := coreService.Run(); err != nil {
		log.Fatal(ctx, "service shutdown: ", err.Error())
	}
}
