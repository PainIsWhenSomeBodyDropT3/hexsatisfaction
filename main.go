package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/JesusG2000/hexsatisfaction-model/pg"
	"github.com/JesusG2000/hexsatisfaction/controllers"
	"github.com/JesusG2000/hexsatisfaction/handler"
	_ "github.com/lib/pq"
)

// Where to place it?

const DIALECT = "postgres"
const HOST = "localhost"
const DbPort = "5432"
const USER = "postgres"
const NAME = "postgres"
const PASSWORD = "18051965q"

const serverPort = 8000

func main() {

	ctx := context.Background()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Database
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", HOST, USER, NAME, PASSWORD, DbPort)
	db, err := sql.Open(DIALECT, dbURI)
	if err != nil {
		log.Fatal(err)
	}

	userDb := pg.NewUserRepository(db)

	// Services
	userService := controllers.NewUser(userDb)

	router := handler.NewRouter(userService)

	coreService := &http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: router,
	}

	go startService(ctx, coreService)

	<-stop
	log.Printf("shutting down server...")
}

func startService(ctx context.Context, coreService *http.Server) {
	if err := coreService.ListenAndServe(); err != nil {
		log.Fatal(ctx, "service shutdown", err.Error())
	}
}
