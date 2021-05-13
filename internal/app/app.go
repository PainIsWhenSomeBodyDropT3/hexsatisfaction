package app

import (
	"context"
	"log"
	"net/http"
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
	"github.com/go-openapi/runtime/middleware"
)

// Run runs hexsatisfaction service
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

	router := handler.NewHandler(services, tokenManager)

	routeSwagger(router)

	srv := server.NewServer(cfg, router)

	go startService(ctx, srv)

	<-stop

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Printf("failed to stop server: %v", err)
	}

	log.Printf("shutting down server...")
}

func startService(ctx context.Context, coreService *server.Server) {
	if err := coreService.Run(); err != nil {
		log.Fatal(ctx, "service shutdown: ", err.Error())
	}
}
func routeSwagger(router *handler.API) {
	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)

	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
}
