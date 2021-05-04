package service

import (
	"log"
	"os"
	"strings"

	"github.com/JesusG2000/hexsatisfaction/internal/config"
	"github.com/JesusG2000/hexsatisfaction/internal/repository"
	"github.com/JesusG2000/hexsatisfaction/pkg/auth"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type api struct {
	*Services
	auth.TokenManager
}

const configPath = "config/main"

func InitTest4Mock() (*api, error) {
	env := ".env"
	envPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	envPath = strings.SplitAfter(envPath, "hexsatisfaction")[0]
	if err := godotenv.Load(envPath + "/" + env); err != nil {
		log.Fatal("Error loading .env file")
	}

	configPath := strings.Split(configPath, "/")

	viper.AddConfigPath(envPath + "/" + configPath[0])
	viper.SetConfigName(configPath[1])
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error read .config file")
	}

	return initServices4Test(), nil
}

func initServices4Test() *api {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal("Init config error: ", err)
	}
	_, repos, err := repository.Connect2Repositories()
	if err != nil {
		return nil
	}

	tokenManager, err := auth.NewManager(cfg.Auth.SigningKey)
	if err != nil {
		log.Fatal(err)
	}

	return &api{
		Services: NewServices(Deps{
			Repos:        repos,
			TokenManager: tokenManager,
		}),
		TokenManager: tokenManager,
	}
}
