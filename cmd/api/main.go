package main

import (
	"net/http"

	"github.com/artyom-kalman/simple-registration-youteam/configs"
	"github.com/artyom-kalman/simple-registration-youteam/internal/database"
	"github.com/artyom-kalman/simple-registration-youteam/internal/handlers"
	"github.com/artyom-kalman/simple-registration-youteam/pkg/logger"
)

func main() {
	logger.InitLogger()

	err := configs.LoadConfig()
	if err != nil {
		logger.Error("error loading config: %v", err)
		return
	}

	port, err := configs.GetEnv("PORT")
	if err != nil {
		logger.Error("error getting port: %v", err)
		return
	}

	dbConfig, err := configs.GetDBConfig()
	if err != nil {
		logger.Error("error getting db config: %v", err)
		return
	}

	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		logger.Error("error connecting to database: %v", err)
		return
	}

	err = db.RunMigration()
	if err != nil {
		logger.Error("error running migration: %v", err)
		return
	}

	handlers.InitHandlers(db)

	http.HandleFunc("/api/user", handlers.HandleNewUser)

	logger.Info("Starting server on port %s", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		logger.Error("error starting server: %v", err)
		return
	}
}
