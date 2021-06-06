package service

import (
	"log"
	"os"
	"strings"

	"github.com/JesusG2000/hexsatisfaction/internal/config"
	"github.com/JesusG2000/hexsatisfaction/internal/repository"
	"github.com/JesusG2000/hexsatisfaction/pkg/auth"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// TestAPI represents a struct for tests api.
type TestAPI struct {
	*Services
	auth.TokenManager
}

const configPath = "config/main"

// InitTest4Mock initialize an a TestAPI for mock testing.
func InitTest4Mock() (*TestAPI, error) {
	envPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	envPath = strings.SplitAfter(envPath, "hexsatisfaction")[0]
	if err := godotenv.Load(envPath + "/" + ".env"); err != nil {
		return nil, errors.Wrap(err, "couldn't load env file")
	}

	configPath := strings.Split(configPath, "/")

	viper.AddConfigPath(envPath + "/" + configPath[0])
	viper.SetConfigName(configPath[1])
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "couldn't read config")
	}

	return initServices4Test()
}

func initServices4Test() (*TestAPI, error) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal("Init config error: ", err)
	}
	_, repos, err := repository.Connect2Repositories()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't connect to db")
	}

	tokenManager, err := auth.NewManager(cfg.Auth.SigningKey)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't create jwt manager")
	}

	return &TestAPI{
		Services: NewServices(Deps{
			Repos:        repos,
			TokenManager: tokenManager,
		}),
		TokenManager: tokenManager,
	}, nil
}
