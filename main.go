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
	"github.com/JesusG2000/hexsatisfaction/repository/pg"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {

	ctx := context.Background()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database
	factory := getFactory()

	userDb := factory.NewUserRepository()

	// Services
	userService := controller.NewUser(userDb)

	router := handler.NewRouter(userService)

	coreService := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("port")),
		Handler: router,
	}

	go startService(ctx, coreService)

	<-stop
	log.Printf("shutting down server...")
}

func getFactory() *pg.Factory {
	f, err := pg.NewFactory()
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func startService(ctx context.Context, coreService *http.Server) {
	if err := coreService.ListenAndServe(); err != nil {
		log.Fatal(ctx, "service shutdown", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
