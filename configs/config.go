package configs

import (
	"fmt"
	"github.com/miladvatankhah/go-maker-checker/internal/message_approval/infrastructure/transport/http"
	"github.com/miladvatankhah/go-maker-checker/pkg/clients/postgres"
	"github.com/miladvatankhah/go-maker-checker/pkg/clients/rabbit"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

type Config struct {
	Server   http.Config `mapstructure:"server"`
	Postgres postgres.Config
	Rabbit   rabbit.Config `mapstructure:"rabbit"`
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		return nil, err
	}

	cfg.Postgres = postgres.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     port,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	rabbitUrl := fmt.Sprintf("amqp://%s:%s@%s:%s%s", os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_USER"),
		os.Getenv("RABBIT_HOST"), os.Getenv("RABBIT_PORT"), os.Getenv("RABBIT_VHOST"))

	cfg.Rabbit.Url = rabbitUrl

	return &cfg, nil
}
