package configs

import (
	"errors"
	"os"

	"github.com/artyom-kalman/simple-registration-youteam/pkg/logger"
)

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func GetDBConfig() (*DBConfig, error) {
	if !isConfigLoaded {
		return nil, errors.New("error loading database configuration")
	}

	logger.Info("Loaded database configuration")
	return &DBConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Name:     os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}, nil
}
