package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/JesusG2000/hexsatisfaction/controller"
	"github.com/JesusG2000/hexsatisfaction/handler"
	"github.com/JesusG2000/hexsatisfaction/jwt"
	"github.com/JesusG2000/hexsatisfaction/repository/pg"
	"github.com/spf13/viper"
)

func main() {

	ctx := context.Background()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Database
	pgRepository := PgRepository()

	//JWT
	tokenManager := JWTManager()

	userDb := pgRepository.NewUserRepository()

	// Services
	userService := controller.NewUser(userDb, tokenManager)

	router := handler.NewRouter(userService, tokenManager)

	coreService := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("port")),
		Handler: router,
	}

	go startService(ctx, coreService)

	<-stop
	log.Printf("shutting down server...")
}

func PgRepository() *pg.Repository {
	f, err := pg.NewPgRepository()
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func JWTManager() *jwt.Manager {
	secret := viper.GetString("auth.secret")
	m, err := jwt.NewManager(secret)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func startService(ctx context.Context, coreService *http.Server) {
	if err := coreService.ListenAndServe(); err != nil {
		log.Fatal(ctx, "service shutdown", err.Error())
	}
}
