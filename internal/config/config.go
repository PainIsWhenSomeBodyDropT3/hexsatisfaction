package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type (
	// Config represents a structure with configs for this microservice.
	Config struct {
		Pg   PgConfig
		Auth JWTConfig
		HTTP HTTPConfig
		GRPC GRPCConfig
	}
	// PgConfig represents a structure with configs for pg database.
	PgConfig struct {
		URI             string
		User            string `required:"true"`
		Password        string `required:"true"`
		Host            string `required:"true"`
		Port            int    `required:"true"`
		DatabaseName    string `split_words:"true" required:"true"`
		DatabaseSslMode string `split_words:"true" required:"true"`
		DatabaseDialect string `split_words:"true" required:"true"`
	}
	// JWTConfig represents a structure with configs for jwt-token.
	JWTConfig struct {
		SigningKey string `split_words:"true" required:"true"`
	}
	// HTTPConfig represents a structure with configs for http server.
	HTTPConfig struct {
		Port           int           `required:"true"`
		MaxHeaderBytes int           `split_words:"true" required:"true"`
		ReadTimeout    time.Duration `split_words:"true" required:"true"`
		WriteTimeout   time.Duration `split_words:"true" required:"true"`
	}
	// GRPCConfig represents a structure with configs for grpc.
	GRPCConfig struct {
		Host string `required:"true"`
		Port string `required:"true"`
	}
)

const (
	PG   = "PG"
	JWT  = "JWT"
	HTTP = "HTTP"
	GRPC = "GRPC"
)

// Init populates Config struct with values.
func Init() (*Config, error) {
	var cfg Config

	if err := envconfig.Process(PG, &cfg.Pg); err != nil {
		return nil, errors.Wrap(err, "couldn't process pg")
	}

	if err := envconfig.Process(JWT, &cfg.Auth); err != nil {
		return nil, errors.Wrap(err, "couldn't process jwt")
	}

	if err := envconfig.Process(HTTP, &cfg.HTTP); err != nil {
		return nil, errors.Wrap(err, "couldn't process http")
	}

	if err := envconfig.Process(GRPC, &cfg.GRPC); err != nil {
		return nil, errors.Wrap(err, "couldn't process grpc")
	}

	return &cfg, nil
}
