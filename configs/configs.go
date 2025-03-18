package configs

import (
	"errors"
	"fmt"
	"os"

	"github.com/artyom-kalman/simple-registration-youteam/pkg/logger"

	"github.com/joho/godotenv"
)

var isConfigLoaded bool = false

func LoadConfig() error {
	if _, err := os.Stat(".env"); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}
	isConfigLoaded = true

	logger.Info("Loaded .env file")

	return nil
}

func GetEnv(key string) (string, error) {
	if !isConfigLoaded {
		return "", errors.New("error loading env variable")
	}

	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("Environment variable %s is not set", key)
	}

	return value, nil
}
