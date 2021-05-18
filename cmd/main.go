package main

import "github.com/JesusG2000/hexsatisfaction/internal/app"

const configPath = "config/main"

// @title Hexsatisfaction API
// @version 1.0
// @description API Service for Hexsatisfaction

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	app.Run(configPath)
}
