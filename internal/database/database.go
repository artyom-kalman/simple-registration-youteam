package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/artyom-kalman/simple-registration-youteam/configs"
	"github.com/artyom-kalman/simple-registration-youteam/pkg/logger"

	_ "github.com/lib/pq"
)

var (
	maxRetries = 5
	retryDelay = 5 * time.Second
)

type Database struct {
	conn *sql.DB
}

func NewDatabase(config *configs.DBConfig) (*Database, error) {
	logger.Info("Initializing database connection")

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name,
	)
	logger.Debug("Connection string: %s", connString)

	var db *sql.DB
	var err error
	for range maxRetries {
		db, err = sql.Open("postgres", connString)
		if err = db.Ping(); err != nil {
			logger.Info("Failed to connect: %v. Retrying...", err)
			time.Sleep(retryDelay)
			continue
		}
		break
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connection to database: %v", err)
	}

	return &Database{
		conn: db,
	}, nil
}
