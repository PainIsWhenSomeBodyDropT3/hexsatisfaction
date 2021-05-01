package test_util

import (
	"log"
	"os"
	"strings"

	"github.com/JesusG2000/hexsatisfaction/jwt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type api struct {
	*jwt.Manager
}

func InitTest4Mock() (*api, error) {
	env := ".env"
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path = strings.SplitAfter(path, "hexsatisfaction")[0]
	if err := godotenv.Load(path + "/" + env); err != nil {
		log.Fatal("Error loading .env file")
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error read .env file")
	}

	return initTestApi(), nil
}

func initTestApi() *api {
	secret := viper.GetString("auth.secret")
	tokenManager, err := jwt.NewManager(secret)
	if err != nil {
		log.Fatal(err)
	}
	return &api{tokenManager}
}
