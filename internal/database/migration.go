package database

import (
	"fmt"

	"github.com/artyom-kalman/simple-registration-youteam/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (db *Database) RunMigration() error {
	logger.Info("Running migration...")

	driver, err := postgres.WithInstance(db.conn, &postgres.Config{})
	if err != nil {
		return err
	}

	migration, err := migrate.NewWithDatabaseInstance(
		"file://migrations/",
		"youteam",
		driver,
	)
	if err != nil {
		return err
	}

	err = migration.Up()
	if err == migrate.ErrNoChange {
		logger.Info("Database schema is up to date")
		return nil
	} else if err != nil {
		return fmt.Errorf("Migration error: %v", err)
	}

	logger.Info("Migration completed successfully")
	return nil

}
