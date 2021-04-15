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
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	ctx := context.Background()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	dialect, _ := os.LookupEnv("DIALECT")
	host, _ := os.LookupEnv("HOST")
	user, _ := os.LookupEnv("DB_USER")
	name, _ := os.LookupEnv("NAME")
	password, _ := os.LookupEnv("PASSWORD")
	dbPort, _ := os.LookupEnv("DB_PORT")
	serverPort, _ := os.LookupEnv("SERVER_PORT")

	// Database
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, name, password, dbPort)
	log.Println(dbURI)

	db, err := sql.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	}

	userDb := pg.NewUserRepository(db)
	// Services
	userService := controllers.NewUser(userDb)

	router := handler.NewRouter(userService)

	coreService := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
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
