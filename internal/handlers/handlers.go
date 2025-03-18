package handlers

import (
	"github.com/artyom-kalman/simple-registration-youteam/internal/database"
)

var db database.Database

func InitHandlers(database *database.Database) {
	db = *database
}
