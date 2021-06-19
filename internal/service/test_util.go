package service

import (
	"log"

	"github.com/JesusG2000/hexsatisfaction/internal/config"
	"github.com/JesusG2000/hexsatisfaction/internal/repository"
	"github.com/JesusG2000/hexsatisfaction/pkg/auth"
	"github.com/pkg/errors"
)

// TestAPI represents a struct for tests api.
type TestAPI struct {
	*Services
	auth.TokenManager
}

// InitTest4Mock initialize an a TestAPI for mock testing.
func InitTest4Mock() (*TestAPI, error) {
	test, err := initServices4Test()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't init tests for mock")
	}

	return test, nil
}

func initServices4Test() (*TestAPI, error) {
	cfg, err := config.Init()
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
