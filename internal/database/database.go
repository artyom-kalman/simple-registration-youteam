package database

import (
	"database/sql"
	"fmt"

	"github.com/artyom-kalman/simple-registration-youteam/configs"
	"github.com/artyom-kalman/simple-registration-youteam/pkg/logger"

	_ "github.com/lib/pq"
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

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error connection to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connection to database: %v", err)
	}

	return &Database{
		conn: db,
	}, nil
}
