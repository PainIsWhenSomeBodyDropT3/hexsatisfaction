package config

import (
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type (
	// Config represents a structure with configs for this microservice.
	Config struct {
		Pg   PgConfig
		Auth JWTConfig
		HTTP HTTPConfig
	}
	// PgConfig represents a structure with configs for pg database.
	PgConfig struct {
		URI      string
		User     string
		Password string
		Host     string
		Port     int
		Name     string
		SslMode  string
		Dialect  string
	}
	// JWTConfig represents a structure with configs for jwt-token.
	JWTConfig struct {
		SigningKey string
	}
	// HTTPConfig represents a structure with configs for http server.
	HTTPConfig struct {
		Port           int
		MaxHeaderBytes int
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
	}
)

// Init populates Config struct with values from config file located at filepath and environment variables.
func Init(path string) (*Config, error) {

	if err := parseConfigFile(path); err != nil {
		return nil, errors.Wrap(err, "couldn't parse config file")
	}

	if err := parseEnv(); err != nil {
		return nil, errors.Wrap(err, "couldn't parse env file")
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "couldn't unmarshal config file")
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func setFromEnv(cfg *Config) {
	cfg.Auth.SigningKey = viper.GetString("signing_key")

	cfg.Pg.User = viper.GetString("user")
	cfg.Pg.Password = viper.GetString("password")
	cfg.Pg.Host = viper.GetString("host")
	cfg.Pg.Port = viper.GetInt("port")
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http.port", &cfg.HTTP.Port); err != nil {
		return errors.Wrap(err, "couldn't unmarshal http.port")
	}

	if err := viper.UnmarshalKey("http.maxHeaderBytes", &cfg.HTTP.MaxHeaderBytes); err != nil {
		return errors.Wrap(err, "couldn't unmarshal http.maxHeaderBytes")
	}

	if err := viper.UnmarshalKey("http.readTimeout", &cfg.HTTP.ReadTimeout); err != nil {
		return errors.Wrap(err, "couldn't unmarshal http.readTimeout")
	}

	if err := viper.UnmarshalKey("http.writeTimeout", &cfg.HTTP.WriteTimeout); err != nil {
		return errors.Wrap(err, "couldn't unmarshal http.writeTimeout")
	}

	if err := viper.UnmarshalKey("pg.databaseName", &cfg.Pg.Name); err != nil {
		return errors.Wrap(err, "couldn't unmarshal pg.databaseName")
	}

	if err := viper.UnmarshalKey("pg.databaseSllMode", &cfg.Pg.SslMode); err != nil {
		return errors.Wrap(err, "couldn't unmarshal pg.databaseSllMode")
	}

	if err := viper.UnmarshalKey("pg.databaseDialect", &cfg.Pg.SslMode); err != nil {
		return errors.Wrap(err, "couldn't unmarshal pg.databaseDialect")
	}

	return nil
}

func parseConfigFile(filepath string) error {
	envPath, err := os.Getwd()
	if err != nil {
		return err
	}
	envPath = strings.SplitAfter(envPath, "hexsatisfaction")[0]
	if err := godotenv.Load(envPath + "/" + ".env"); err != nil {
		return errors.Wrap(err, "couldn't load env file")
	}

	configPath := strings.Split(filepath, "/")

	viper.AddConfigPath(envPath + "/" + configPath[0])
	viper.SetConfigName(configPath[1])
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "couldn't load config file")
	}

	return nil
}

func parseEnv() error {

	if err := parsePg(); err != nil {
		return errors.Wrap(err, "couldn't parse pg")
	}

	if err := parseJWT(); err != nil {
		return errors.Wrap(err, "couldn't parse jwt")
	}

	return nil
}

func parseJWT() error {
	viper.SetEnvPrefix("jwt")

	if err := viper.BindEnv("signing_key"); err != nil {
		return errors.Wrap(err, "couldn't bind signing_key")
	}

	return nil
}

func parsePg() error {
	viper.SetEnvPrefix("pg")

	if err := viper.BindEnv("password"); err != nil {
		return errors.Wrap(err, "couldn't bind password")
	}

	if err := viper.BindEnv("user"); err != nil {
		return errors.Wrap(err, "couldn't bind user")
	}

	if err := viper.BindEnv("host"); err != nil {
		return errors.Wrap(err, "couldn't bind host")
	}

	if err := viper.BindEnv("port"); err != nil {
		return errors.Wrap(err, "couldn't bind port")
	}

	return nil
}
